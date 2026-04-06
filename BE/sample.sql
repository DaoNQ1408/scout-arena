-- 1. Định nghĩa ENUM
CREATE TYPE user_rank AS ENUM ('BRONZE', 'SILVER', 'GOLD', 'PLATINUM');
CREATE TYPE participation_status AS ENUM ('NOT_STARTED', 'PARTICIPATED', 'COMPLETED');

-- 2. Định nghĩa cấu trúc giải đấu
CREATE TABLE seasons (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(255) NOT NULL,
                         is_active BOOLEAN DEFAULT TRUE
);

CREATE TABLE rounds (
                        id SERIAL PRIMARY KEY,
                        season_id INTEGER REFERENCES seasons(id),
                        name VARCHAR(255) NOT NULL,
                        base_points INTEGER DEFAULT 100, -- Điểm tối đa của cả Round
                        "order" INTEGER
);

CREATE TABLE challenges (
                            id SERIAL PRIMARY KEY,
                            round_id INTEGER REFERENCES rounds(id),
                            rank_level user_rank NOT NULL,
                            title VARCHAR(255) NOT NULL,
                            weight_percentage DECIMAL(5, 2) NOT NULL -- % điểm của challenge này trong Round
);

-- 3. Người dùng và Tiến trình
CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       username VARCHAR(100) UNIQUE NOT NULL,
                       current_rank user_rank DEFAULT 'BRONZE'
);

-- Theo dõi Round: Lưu rank lúc bắt đầu để không bị thay đổi challenge khi đang làm
CREATE TABLE user_round_progress (
                                     user_id INTEGER REFERENCES users(id),
                                     round_id INTEGER REFERENCES rounds(id),
                                     rank_at_start user_rank NOT NULL,
                                     total_round_points INTEGER DEFAULT 0, -- Tổng điểm đã tích lũy trong Round này
                                     PRIMARY KEY (user_id, round_id)
);

-- Theo dõi từng Challenge: Tính điểm ngay khi status = 'COMPLETED'
CREATE TABLE user_challenge_records (
                                        id SERIAL PRIMARY KEY,
                                        user_id INTEGER REFERENCES users(id),
                                        challenge_id INTEGER REFERENCES challenges(id),
                                        status participation_status DEFAULT 'NOT_STARTED',
                                        earned_points INTEGER DEFAULT 0, -- Điểm nhận được cho riêng challenge này
                                        completed_at TIMESTAMP,
                                        UNIQUE(user_id, challenge_id)
);

-- Tổng điểm mùa giải
CREATE TABLE user_season_stats (
                                   user_id INTEGER REFERENCES users(id),
                                   season_id INTEGER REFERENCES seasons(id),
                                   total_season_points INTEGER DEFAULT 0,
                                   PRIMARY KEY (user_id, season_id)
);
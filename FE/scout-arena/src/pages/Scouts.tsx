import type { Scout } from "@/types/scout";
import ScoutCard from "@/components/shared/ScoutCard";

const MOCK_DRIVERS: Scout[] = [
  {
    id: "1",
    firstName: "George",
    lastName: "Russell",
    team: "Mercedes",
    number: 63,
    countryFlag: "🇬🇧",
    image: "/images/russell.png", // Replace with your actual asset path
    brandColor: "#27F4D2",
  },
  {
    id: "2",
    firstName: "Charles",
    lastName: "Leclerc",
    team: "Ferrari",
    number: 16,
    countryFlag: "🇲🇨",
    image: "/images/leclerc.png",
    brandColor: "#E80020",
  },
  {
    id: "3",
    firstName: "George",
    lastName: "Russell",
    team: "Mercedes",
    number: 63,
    countryFlag: "🇬🇧",
    image: "/images/russell.png", // Replace with your actual asset path
    brandColor: "#27F4D2",
  },
  {
    id: "4",
    firstName: "Charles",
    lastName: "Leclerc",
    team: "Ferrari",
    number: 16,
    countryFlag: "🇲🇨",
    image: "/images/leclerc.png",
    brandColor: "#E80020",
  },
  {
    id: "5",
    firstName: "George",
    lastName: "Russell",
    team: "Mercedes",
    number: 63,
    countryFlag: "🇬🇧",
    image: "/images/russell.png", // Replace with your actual asset path
    brandColor: "#27F4D2",
  },
  {
    id: "6",
    firstName: "Charles",
    lastName: "Leclerc",
    team: "Ferrari",
    number: 16,
    countryFlag: "🇲🇨",
    image: "/images/leclerc.png",
    brandColor: "#E80020",
  },
  // ... add more drivers
];

export default function Scouts() {
  return (
    <div className="container mx-auto py-10 px-4">
      <header className="mb-10">
        <h1 className="text-4xl font-black uppercase italic border-b-4 border-red-600 inline-block">
          F1 Drivers 2026
        </h1>
      </header>

      {/* 4-column Grid: Responsive for mobile/tablet */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
        {MOCK_DRIVERS.map((driver) => (
          <ScoutCard key={driver.id} scout={driver} />
        ))}
      </div>
    </div>
  );
}

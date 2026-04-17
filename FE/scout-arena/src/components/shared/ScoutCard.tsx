import { Card } from "@/components/ui/card";
import type { Scout } from "@/types/scout";

interface ScoutCardProps {
  scout: Scout;
}

export default function ScoutCard({ scout }: ScoutCardProps) {
  return (
    <Card className="group relative h-[300px] w-full overflow-hidden border-none rounded-xl transition-transform hover:-translate-y-2 duration-300 cursor-pointer shadow-lg">
      {/* Dynamic Background Color based on Team */}
      <div
        className="absolute inset-0 opacity-90 transition-opacity group-hover:opacity-100"
        style={{ backgroundColor: scout.brandColor }}
      />

      {/* Card Content */}
      <div className="relative z-10 p-5 flex flex-col h-full text-white">
        <div className="flex justify-between items-start">
          <div>
            <p className="text-sm font-medium opacity-90">{scout.firstName}</p>
            <h3 className="text-2xl font-black uppercase leading-none tracking-tighter">
              {scout.lastName}
            </h3>
            <p className="text-xs mt-1 font-semibold opacity-75">
              {scout.team}
            </p>
          </div>
          <span className="text-4xl font-black italic opacity-40">
            {scout.number}
          </span>
        </div>

        {/* Driver Image */}
        <div className="mt-auto self-end w-full flex justify-end overflow-hidden">
          <img
            src={scout.image}
            alt={scout.lastName}
            className="h-48 object-contain transition-transform duration-500 group-hover:scale-110"
          />
        </div>

        {/* Footer/Flag */}
        <div className="absolute bottom-4 left-5">
          <span className="text-2xl">{scout.countryFlag}</span>
        </div>
      </div>
    </Card>
  );
}

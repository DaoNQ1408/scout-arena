import { Search, Home, Pencil, GitBranch, Tag, Trash2 } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Badge } from "@/components/ui/badge";

/* ───────── Color Palette Card ───────── */
function ColorCard({
  name,
  hex,
  shades,
  bgClass,
  textClass = "text-white",
}: {
  name: string;
  hex: string;
  shades: string[];
  bgClass: string;
  textClass?: string;
}) {
  return (
    <div className="rounded-xl overflow-hidden">
      <div className={`${bgClass} px-4 pt-4 pb-3 flex justify-between items-start ${textClass}`}>
        <span className="font-heading font-semibold text-sm">{name}</span>
        <span className="font-mono text-xs opacity-80">{hex}</span>
      </div>
      <div className="grid grid-cols-5 gap-[2px] p-1 bg-card rounded-b-xl">
        {shades.map((s, i) => (
          <div
            key={i}
            className="aspect-square rounded-md"
            style={{ backgroundColor: s }}
          />
        ))}
      </div>
    </div>
  );
}

/* ───────── Glass Card wrapper ───────── */
function GlassCard({
  children,
  className = "",
}: {
  children: React.ReactNode;
  className?: string;
}) {
  return (
    <div
      className={`rounded-2xl bg-card border border-border p-5 ${className}`}
    >
      {children}
    </div>
  );
}

/* ───────── Main Page ───────── */
export default function DesignSystemPage() {
  /* palette shades */
  const primaryShades = [
    "#E10600", "#E8382E", "#F06A5C", "#F5A59E", "#FCDAD8",
    "#FDE8E7", "#FCF0EF", "#FEF7F6", "#FFFBFA", "#FFF5F5",
  ];
  const secondaryShades = [
    "#6C727D", "#838892", "#9BA0A8", "#B3B6BD", "#CBCDD2",
    "#DEE0E3", "#EAEBEE", "#F2F3F5", "#F8F8F9", "#FCFCFD",
  ];
  const tertiaryShades = [
    "#FFFFFF", "#FAFAFA", "#F5F5F5", "#E8E8E8", "#D4D4D4",
    "#C5C5C5", "#B0B0B0", "#A0A0A0", "#8A8A8A", "#757575",
  ];
  const neutralShades = [
    "#1A1A1D", "#2A2A2E", "#3A3A3E", "#4A4A4E", "#5A5A5E",
    "#6A6A6E", "#7A7A7E", "#8A8A8E", "#9A9A9E", "#AAAAAE",
  ];

  return (
    <div className="min-h-screen bg-background p-4 md:p-8">
      <div className="max-w-[1200px] mx-auto grid grid-cols-12 gap-4 auto-rows-min">

        {/* ─── LEFT COLUMN: Color Palette ─── */}
        <div className="col-span-12 md:col-span-3 flex flex-col gap-4">
          <ColorCard
            name="Primary"
            hex="#E10600"
            bgClass="bg-brand-primary"
            shades={primaryShades}
          />
          <ColorCard
            name="Secondary"
            hex="#6C727D"
            bgClass="bg-brand-secondary"
            shades={secondaryShades}
          />
          <ColorCard
            name="Tertiary"
            hex="#FFFFFF"
            bgClass="bg-white"
            textClass="text-neutral-800"
            shades={tertiaryShades}
          />
          <ColorCard
            name="Neutral"
            hex="#1A1A1D"
            bgClass="bg-brand-neutral"
            shades={neutralShades}
          />
        </div>

        {/* ─── RIGHT AREA: 3×3 bento grid ─── */}
        <div className="col-span-12 md:col-span-9 grid grid-cols-3 gap-4 auto-rows-min">

          {/* Row 1 ─────────────────────────── */}
          {/* Typography – Headline */}
          <GlassCard>
            <div className="flex justify-between items-center mb-3">
              <span className="text-muted-foreground text-xs tracking-widest uppercase">Headline</span>
              <span className="text-muted-foreground text-xs">Space Grotesk</span>
            </div>
            <p className="font-heading text-6xl font-medium text-[#E8A09A]">Aa</p>
          </GlassCard>

          {/* Buttons */}
          <GlassCard>
            <div className="grid grid-cols-2 gap-2">
              <Button className="bg-brand-primary hover:bg-brand-primary/90 text-white rounded-md text-sm h-9">
                Primary
              </Button>
              <Button
                variant="secondary"
                className="bg-brand-secondary hover:bg-brand-secondary/80 text-white rounded-md text-sm h-9"
              >
                Secondary
              </Button>
              <Button
                variant="outline"
                className="border-brand-primary/40 text-brand-primary hover:bg-brand-primary/10 rounded-md text-sm h-9"
              >
                Inverted
              </Button>
              <Button
                variant="outline"
                className="border-brand-primary text-brand-primary hover:bg-brand-primary/10 rounded-md text-sm h-9"
              >
                Outlined
              </Button>
            </div>
          </GlassCard>

          {/* Search */}
          <GlassCard>
            <div className="relative">
              <Search className="absolute left-3 top-1/2 -translate-y-1/2 h-4 w-4 text-muted-foreground" />
              <Input
                placeholder="Search"
                className="pl-9 bg-transparent border-brand-secondary/40 text-foreground placeholder:text-muted-foreground focus-visible:ring-brand-primary/50 rounded-md"
              />
            </div>
          </GlassCard>

          {/* Row 2 ─────────────────────────── */}
          {/* Typography – Body */}
          <GlassCard>
            <div className="flex justify-between items-center mb-3">
              <span className="text-muted-foreground text-xs tracking-widest uppercase">Body</span>
              <span className="text-muted-foreground text-xs">Inter</span>
            </div>
            <p className="font-sans text-6xl font-medium text-[#E8A09A]">Aa</p>
          </GlassCard>

          {/* Lines / separators */}
          <GlassCard className="flex flex-col justify-center gap-3">
            <div className="h-[3px] rounded-full bg-brand-primary" />
            <div className="h-[3px] rounded-full bg-brand-primary/70" />
            <div className="h-[2px] rounded-full bg-brand-secondary" />
            <div className="h-[2px] rounded-full bg-brand-secondary/50" />
          </GlassCard>

          {/* Icon Nav */}
          <GlassCard className="flex items-center justify-center">
            <div className="flex gap-2">
              <button className="w-10 h-10 rounded-lg bg-brand-primary/20 border border-brand-primary/40 flex items-center justify-center text-brand-primary hover:bg-brand-primary/30 transition-colors">
                <Home className="w-4 h-4" />
              </button>
              <button className="w-10 h-10 rounded-lg bg-transparent border border-border flex items-center justify-center text-muted-foreground hover:text-foreground hover:border-brand-primary/40 transition-colors">
                <Search className="w-4 h-4" />
              </button>
              <button className="w-10 h-10 rounded-lg bg-transparent border border-border flex items-center justify-center text-muted-foreground hover:text-foreground hover:border-brand-primary/40 transition-colors">
                <svg className="w-4 h-4" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                  <path d="M20 21v-2a4 4 0 0 0-4-4H8a4 4 0 0 0-4 4v2" />
                  <circle cx="12" cy="7" r="4" />
                </svg>
              </button>
            </div>
          </GlassCard>

          {/* Row 3 ─────────────────────────── */}
          {/* Typography – Label */}
          <GlassCard>
            <div className="flex justify-between items-center mb-3">
              <span className="text-muted-foreground text-xs tracking-widest uppercase">Label</span>
              <span className="text-muted-foreground text-xs">Space Grotesk</span>
            </div>
            <p className="font-heading text-6xl font-medium text-[#E8A09A]">Aa</p>
          </GlassCard>

          {/* Edit / Label badges */}
          <GlassCard className="flex items-center justify-center gap-3">
            <div className="w-10 h-10 rounded-lg bg-muted border border-border flex items-center justify-center text-muted-foreground">
              <Pencil className="w-4 h-4" />
            </div>
            <Badge className="bg-brand-primary hover:bg-brand-primary/90 text-white px-4 py-1.5 text-sm rounded-md flex items-center gap-1.5">
              <Pencil className="w-3 h-3" />
              Label
            </Badge>
          </GlassCard>

          {/* Action icons */}
          <GlassCard className="flex items-center justify-center">
            <div className="flex gap-2">
              <button className="w-10 h-10 rounded-lg bg-brand-primary/20 border border-brand-primary/40 flex items-center justify-center text-brand-primary hover:bg-brand-primary/30 transition-colors">
                <Pencil className="w-4 h-4" />
              </button>
              <button className="w-10 h-10 rounded-lg bg-brand-primary/20 border border-brand-primary/40 flex items-center justify-center text-brand-primary hover:bg-brand-primary/30 transition-colors">
                <GitBranch className="w-4 h-4" />
              </button>
              <button className="w-10 h-10 rounded-lg bg-brand-primary/20 border border-brand-primary/40 flex items-center justify-center text-brand-primary hover:bg-brand-primary/30 transition-colors">
                <Tag className="w-4 h-4" />
              </button>
              <button className="w-10 h-10 rounded-lg bg-brand-primary/20 border border-brand-primary/40 flex items-center justify-center text-brand-primary hover:bg-brand-primary/30 transition-colors">
                <Trash2 className="w-4 h-4" />
              </button>
            </div>
          </GlassCard>

        </div>
      </div>
    </div>
  );
}

"use client";

import { FacilityResp } from "@/lib/fetcher";
import { Card } from "./card";
import Image from "next/image";
import { cn } from "@/lib/utils";
import { Typography as typography } from "./typography";
import { Clock, Star } from "lucide-react";
import { Badge, badgeVariants } from "./badge";
import { Button } from "./button";

type StandbyCardProps = {
  facility: FacilityResp;
  size: "sm" | "lg";
  showImage?: boolean;
  onFavorite: (facilityId: string) => void;
  isFavorite: boolean;
  imageUrl?: string;
};

export const StandbyCard = ({
  facility,
  size,
  showImage = false,
  onFavorite,
  isFavorite,
  imageUrl,
}: StandbyCardProps) => {
  if (size === "sm") {
    return (
      <SmallStandbyCard
        facility={facility}
        showImage={showImage}
        onFavorite={onFavorite}
        isFavorite={isFavorite}
        imageUrl={imageUrl}
      />
    );
  }

  return <Card>StandbyCard</Card>;
};

const SmallStandbyCard = ({
  facility,
  showImage,
  imageUrl,
  isFavorite,
  onFavorite,
}: {
  facility: FacilityResp;
  showImage: boolean;
  imageUrl?: string;
  isFavorite: boolean;
  onFavorite: (facilityId: string) => void;
}) => {
  const defaultImage =
    "https://images.unsplash.com/photo-1588345921523-c2dcdb7f1dcd?w=800&dpr=2&q=80";
  return (
    <Card className="flex flex-row gap-4 p-0">
      <div className="relative w-[40%] max-w-[120px]">
        <Image
          src={imageUrl || defaultImage}
          alt="Photo by Drew Beamer"
          fill
          className="rounded-md rounded-r-none object-cover"
        />
      </div>
      <div className="mr-4 flex w-full flex-col gap-2 py-2">
        <h1
          className={cn(
            typography({ variant: "large" }),
            "line-clamp-2 w-full overflow-hidden text-ellipsis",
          )}
        >
          {facility.name}
        </h1>
        <section className="flex gap-1">
          <Button
            variant="link"
            className={cn(
              badgeVariants({ variant: "outline" }),
              "flex h-6 items-center gap-1",
            )}
            onClick={() => onFavorite(facility.id)}
          >
            {isFavorite ? (
              <Star className="h-3 w-3 fill-accent-foreground" />
            ) : (
              <Star className="h-3 w-3" />
            )}
          </Button>
          <Badge variant="outline" className="flex items-center gap-1">
            <Clock className="h-3 w-3" /> {facility.standbyTime}分待ち
          </Badge>
        </section>
      </div>
    </Card>
  );
};

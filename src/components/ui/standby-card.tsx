"use client";

import { FacilityResp } from "@/lib/fetcher";
import { Card } from "./card";
import { AspectRatio } from "./aspect-ratio";
import Image from "next/image";
import { cn } from "@/lib/utils";
import { Typography as typography } from "./typography";
import { Clock, Star } from "lucide-react";
import { Badge, badgeVariants } from "./badge";
import { useState } from "react";
import { Button } from "./button";

type StandbyCardProps = {
  facility: FacilityResp;
  size: "sm" | "lg";
  showImage?: boolean;
  onFavorite: (facilityId: string) => void;
  isFavorite: boolean;
};

export const StandbyCard = ({
  facility,
  size,
  showImage = false,
  onFavorite,
  isFavorite,
}: StandbyCardProps) => {
  if (size === "sm") {
    return (
      <SmallStandbyCard
        facility={facility}
        showImage={showImage}
        onFavorite={onFavorite}
        isFavorite={isFavorite}
      />
    );
  }

  return <Card>StandbyCard</Card>;
};

const SmallStandbyCard = ({
  facility,
  showImage,
  isFavorite,
  onFavorite,
}: {
  facility: FacilityResp;
  showImage: boolean;
  isFavorite: boolean;
  onFavorite: (facilityId: string) => void;
}) => {
  return (
    <Card className="flex flex-row gap-4 p-0">
      <div className="h-full w-[120px]">
        <AspectRatio ratio={16 / 9}>
          <Image
            src="https://images.unsplash.com/photo-1588345921523-c2dcdb7f1dcd?w=800&dpr=2&q=80"
            alt="Photo by Drew Beamer"
            fill
            className="h-full w-full rounded-md rounded-r-none object-cover"
          />
        </AspectRatio>
      </div>
      <div className="mr-4 flex w-full flex-col gap-2 py-2">
        <h1
          className={cn(
            typography({ variant: "large" }),
            "max-h-[56px] flex-1 overflow-hidden text-ellipsis",
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

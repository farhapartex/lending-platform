import type { MetadataRoute } from "next";
import { AppRoute } from "@/lib/enums";
import { absoluteUrl, indexableRoutes } from "@/lib/site";

const priorities: Partial<Record<AppRoute, number>> = {
  [AppRoute.Home]: 1,
  [AppRoute.Markets]: 0.9,
  [AppRoute.Lend]: 0.9,
  [AppRoute.Borrow]: 0.9,
  [AppRoute.Learn]: 0.7,
};

export default function sitemap(): MetadataRoute.Sitemap {
  return indexableRoutes.map((route) => ({
    url: absoluteUrl(route),
    priority: priorities[route] ?? 0.5,
  }));
}

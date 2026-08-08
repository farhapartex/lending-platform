import type { MetadataRoute } from "next";
import { nonIndexableRoutes, siteUrl } from "@/lib/site";

export default function robots(): MetadataRoute.Robots {
  return {
    rules: [
      {
        userAgent: "*",
        allow: "/",
        disallow: nonIndexableRoutes,
      },
    ],
    sitemap: new URL("/sitemap.xml", siteUrl).toString(),
  };
}

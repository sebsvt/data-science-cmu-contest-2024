"use client";

import { ColumnDef } from "@tanstack/react-table";
import { MoreHorizontal, ArrowUpDown } from "lucide-react";

import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import Link from "next/link";

export const columns: ColumnDef<DemandForecastItem>[] = [
  {
    accessorKey: "from_last_date",
    header: "From the last date",
    cell: ({ row }) => {
      const dateString = row.original.from_last_date;
      const formattedDate = new Date(dateString).toLocaleDateString("en-CA"); // Format as YYYY-MM-DD
      return <span className="text-center h-8 w-8">{formattedDate}</span>; // Center the text
    },
  },
  {
    accessorKey: "name",
    header: "Name",
  },
  // New column for next day's date
  {
    accessorKey: "next_day_date",
    header: "Next day date",
    cell: ({ row }) => {
      const predictions: Prediction[] = row.original.predictions;
      const nextDayDate = predictions.length
        ? new Date(predictions[0].future_date).toLocaleDateString("en-CA")
        : "N/A"; // Format next day date as YYYY-MM-DD
      return <span className="text-center">{nextDayDate}</span>; // Center the text
    },
  },
  // Column for next day's demand
  {
    accessorKey: "next_day_demand",
    header: "Next day demand",
    cell: ({ row }) => {
      const predictions: Prediction[] = row.original.predictions;
      const nextDayDemand = predictions.length
        ? Math.round(predictions[0].demand)
        : "N/A";
      return <span className="text-center">{nextDayDemand}</span>; // Center the text
    },
  },
  {
    id: "actions",
    cell: ({ row }) => {
      const demand_forecast = row.original;

      return (
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="ghost" className="h-8 w-8 p-0">
              <span className="sr-only">Open menu</span>
              <MoreHorizontal className="h-4 w-4" />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent align="end">
            <DropdownMenuLabel>Actions</DropdownMenuLabel>
            <DropdownMenuItem
              onClick={() =>
                navigator.clipboard.writeText(demand_forecast.from_last_date)
              }
            >
              Copy Forecast ID
            </DropdownMenuItem>
            <DropdownMenuItem>
              <Link href={"/"}>View Forecast Result</Link>
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      );
    },
  },
];

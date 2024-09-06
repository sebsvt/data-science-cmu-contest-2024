"use client";

import { ColumnDef } from "@tanstack/react-table";
import { MoreHorizontal, ArrowUpDown, Check, X, Loader2 } from "lucide-react";

import { Button } from "@/components/ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import Link from "next/link";

export const columns: ColumnDef<DemandForecast>[] = [
  {
    accessorKey: "title",
    header: "Title",
  },
  {
    accessorKey: "number_of_item",
    header: "Number of item",
  },
  {
    accessorKey: "created_at",
    header: ({ column }) => {
      return (
        <Button
          variant="ghost"
          onClick={() => column.toggleSorting(column.getIsSorted() === "asc")}
        >
          Created at
          <ArrowUpDown className="ml-2 h-4 w-4" />
        </Button>
      );
    },
  },
  {
    accessorKey: "status",
    header: "Status",
    cell: ({ row }) => {
      const status: string = row.getValue("status"); // Get the status value
      let icon;
      switch (status) {
        case "completed":
          icon = <Check className="h-4 w-4 text-green-600" />;
          break;
        case "failed":
          icon = <X className="h-4 w-4 text-red-600" />;
          break;
        case "pending":
          icon = <Loader2 className="h-4 w-4 animate-spin" />;
          break;
        default:
          icon = null;
      }
      return (
        <div className="flex items-center">
          {icon}
          <span className="ml-2">{status}</span>
        </div>
      );
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
                navigator.clipboard.writeText(demand_forecast.forecast_id)
              }
            >
              Copy Forecast ID
            </DropdownMenuItem>
            <DropdownMenuItem>
              <Link
                href={`/business-suite/software/demand-forecasting/${demand_forecast.partner_id}/detail/${demand_forecast.forecast_id}`}
              >
                View Forecast Result
              </Link>
            </DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>
      );
    },
  },
];

import { DataTable } from "@/components/data-table";
import React from "react";
import { columns } from "./forecast_columns";

async function getData(forecast_id: string): Promise<DemandForecast | null> {
  try {
    const res = await fetch(
      `http://89.213.177.102:8082/api/demand-forecast/${forecast_id}`,
      { cache: "no-cache" }
    );

    // Check if the response is ok (status is in the range of 200-299)
    if (!res.ok) {
      throw new Error(`Failed to fetch data: ${res.status} ${res.statusText}`);
    }

    // Parse the response as JSON
    const demand_forecasting: DemandForecast = await res.json();
    return demand_forecasting;
  } catch (error) {
    console.error("Error fetching demand forecast data:", error);
    return null;
  }
}

interface PageProps {
  params: { forecast_id: string };
}

const Page = async ({ params: { forecast_id } }: PageProps) => {
  const data = await getData(forecast_id);
  if (!data) {
    return <h1>No data.</h1>;
  }
  return (
    <div className="container mx-auto px-4">
      <header className="mt-10">
        <h1 className="text-xl md:text-3xl font-semibold">{data.title}</h1>
        <div className="flex gap-4">
          <span className="text-muted-foreground text-xs">
            Builder: {data.builder_id}
          </span>
        </div>
        <p className="mt-4 text-sm text-stone">{data.description}</p>
      </header>
      <main className="container mx-auto py-10">
        <DataTable
          searchColumn="name"
          columns={columns}
          data={data.demand_forecast_item}
        />
      </main>
    </div>
  );
};

export default Page;

import React from "react";
import { columns } from "./columns";
import { DataTable } from "@/components/data-table";

async function getData(partner_id: string): Promise<DemandForecast[] | null> {
  try {
    const res = await fetch(
      `http://89.213.177.102:8082/api/demand-forecast/${partner_id}/partner`,
      { cache: "no-cache" }
    );

    // Check if the response is ok (status is in the range of 200-299)
    if (!res.ok) {
      throw new Error(`Failed to fetch data: ${res.status} ${res.statusText}`);
    }

    // Parse the response as JSON
    const demand_forecasting: DemandForecast[] = await res.json();
    return demand_forecasting;
  } catch (error) {
    console.error("Error fetching demand forecast data:", error);
    return null;
  }
}

interface PageProps {
  params: { partner_id: string };
}

const Page = async ({ params: { partner_id } }: PageProps) => {
  const data = await getData(partner_id);
  if (!data) {
    return <h1>No data.</h1>;
  }
  return (
    <div className="">
      <header className="mt-6">
        <h1 className="text-xl md:text-3xl font-semibold">Demand Forecast</h1>
        <p className="text-muted-foreground text-sm">
          Leveraging data to forecast market demand and improve business
          outcomes
        </p>
      </header>
      <main className="container mx-auto py-10">
        <DataTable searchColumn="title" columns={columns} data={data} />
      </main>
    </div>
  );
};

export default Page;

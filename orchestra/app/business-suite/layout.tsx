import Sidebar from "@/components/software/sidebar";
import React from "react";

const Layout = ({ children }: { children: React.ReactNode }) => {
  return (
    <div className="flex items-start justify-between">
      <Sidebar partner_id="7568ba12-88a1-41c8-9e30-b9a9fa45440d" />
      <main className="w-full">
        <div className="container mx-auto px-4">{children}</div>
      </main>
    </div>
  );
};

export default Layout;

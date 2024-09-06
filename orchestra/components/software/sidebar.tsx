import Link from "next/link";
import { Command, CommandGroup, CommandItem, CommandList } from "../ui/command";

export default function Sidebar({ partner_id }: { partner_id: string }) {
  const menu_lists = [
    {
      group: "General",
      items: [
        {
          link: "/",
          text: "Dashboard",
        },
      ],
    },
    {
      group: "Sotfware",
      items: [
        {
          link: `/business-suite/software/demand-forecasting/${partner_id}`,
          text: "Demand Forecast",
        },
      ],
    },
  ];
  return (
    <div className="hidden md:flex flex-col gap-4 w-[300px] min-h-screen border-r px-2">
      <div>
        <h2>Nim see seng 1987</h2>
      </div>
      <div className="grow">
        <Command>
          <CommandList className="">
            {menu_lists.map((menu, index) => (
              <CommandGroup heading={menu.group} key={index}>
                {menu.items.map((item, index) => (
                  <CommandItem key={index}>
                    <Link href={item.link}>{item.text}</Link>
                  </CommandItem>
                ))}
              </CommandGroup>
            ))}
          </CommandList>
        </Command>
      </div>
    </div>
  );
}

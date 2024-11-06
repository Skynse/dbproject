import { useSignal } from "@preact/signals";
import Counter from "../islands/Counter.tsx";
import ItemSearch from "../islands/SearchWidget.tsx";
import Dashboard from "../islands/Dashboard.tsx";

export default function Home() {
  const count = useSignal(3);
  const idSearchFilter = useSignal("");
  const nameSearchFilter = useSignal("");
  return (
    <div class="mx-4 my-4">
      <div class="text-2xl font-bold ">Home</div>

      <Dashboard
        count={count}
        idSearchFilter={idSearchFilter}
        nameSearchFilter={nameSearchFilter}
      />
    </div>
  );
}

import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";
import { AccountSettings } from "./components/accountSettings";

function App() {
  return (
    <div className="p-6">
      <h1 className="text-red-500 text-xl font-bold mb-4">Project Manager</h1>
      <Accordion className="border-2 rounded-2xl"type="single" collapsible>
        <AccordionItem value="item-1">
          <AccordionTrigger>Is it accessible?</AccordionTrigger>
          <AccordionContent>
            Yes. It adheres to the WAI-ARIA design pattern.
          </AccordionContent>
        </AccordionItem>
      </Accordion>
      <AccountSettings></AccountSettings>
    </div>
  );
}

export default App;

import {
  Accordion,
  AccordionContent,
  AccordionItem,
  AccordionTrigger,
} from "@/components/ui/accordion";

function App() {
  return (
    <div className="p-6">
      <h1 className="text-red-500 text-xl font-bold mb-4">Project Manager</h1>
      <Accordion type="single" collapsible className="w-full border rounded-lg">
        <AccordionItem value="item-1" className="border-b">
          <AccordionTrigger className="p-4 bg-gray-100 hover:bg-gray-200 transition rounded-t-lg">
            Is it accessible?
          </AccordionTrigger>
          <AccordionContent className="p-4 bg-white">
            Yes. It adheres to the WAI-ARIA design pattern.
          </AccordionContent>
        </AccordionItem>
      </Accordion>
    </div>
  );
}

export default App;

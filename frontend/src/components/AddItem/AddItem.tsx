import { useForm } from "@mantine/form";
import { useState } from "react";
import { Modal, Button, Group, TextInput, Textarea } from "@mantine/core";
import { KeyedMutator } from "swr";
import { ENDPOINT, Todo } from "../../App";

function AddItem({ mutate }: { mutate: KeyedMutator<Todo[]> }) {
  const [open, setOpen] = useState(false);

  const form = useForm({
    initialValues: {
      title: "",
      description: "",
    },
  });

  async function createTodo(values: { title: string; description: string }) {
    try {
      const updated = await fetch(`${ENDPOINT}/api/todos`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(values),
      }).then((r) => r.json());
      mutate(updated);
      form.reset();
      setOpen(false);
      // Handle the response
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  }

  return (
    <>
      <Modal opened={open} onClose={() => setOpen(false)} title="Create todo">
        <form onSubmit={form.onSubmit(createTodo)}>
          <TextInput
            required
            mb={12}
            label="Todo"
            placeholder="What do you want to do?"
            {...form.getInputProps("title")}
          />
          <Textarea
            required
            mb={12}
            label="Body"
            placeholder="Tell me more..."
            {...form.getInputProps("description")}
          />

          <Button type="submit">Create todo</Button>
        </form>
      </Modal>

      <Group position="center">
        <Button onClick={() => setOpen(true)}>New To-do</Button>
      </Group>
    </>
  );
}

export default AddItem;

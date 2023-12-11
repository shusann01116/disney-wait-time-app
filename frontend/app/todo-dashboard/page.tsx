import { generateServerClientUsingCookies } from "@aws-amplify/adapter-nextjs/api";
import { cookies } from "next/headers";
import { revalidatePath } from "next/cache";
import * as mutations from "@/src/graphql/mutations";
import * as queries from "@/src/graphql/queries";
import config from "@/src/amplifyconfiguration.json";

const cookiesClient = generateServerClientUsingCookies({
  config,
  cookies,
});

async function createTodo(formData: FormData) {
  "use server";
  const { data } = await cookiesClient.graphql({
    query: mutations.createTodo,
    variables: {
      input: {
        name: formData.get("name")?.toString() ?? "",
      },
    },
  });

  console.log("Create Todo: ", data?.createTodo);

  revalidatePath("/todo-dashboard");
}

export default async function Home() {
  const { data, errors } = await cookiesClient.graphql({
    query: queries.listTodos,
  });

  const todos = data?.listTodos?.items ?? [];

  return (
    <div>
      <form className="flex items-center" action={createTodo}>
        <input
          name="name"
          placeholder="Add a todo"
          className="p-2 border text-black border-gray-300 rounded-md mr-2"
        />
        <button
          type="submit"
          className="px-4 py-2 bg-blue-500 text-white rounded-md"
        >
          Add
        </button>
      </form>

      {(!todos || todos.length === 0) && errors && (
        <div>
          <p className="mt-4">No todos found</p>
        </div>
      )}

      <ul className="mt-4">
        {todos.map((todo) => (
          <li key={todo?.id} className="p-2 border border-gray-300 rounded-md">
            {todo?.name}
          </li>
        ))}
      </ul>
    </div>
  );
}

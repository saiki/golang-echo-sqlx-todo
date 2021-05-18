import { writable } from 'svelte/store';

function createTodoStore() {
	const { subscribe, set, update } = writable(new Array<Todo>());

	return {
		subscribe,
        load: async () => {
            let list = await fetch("http://localhost:8080/api/todo",
                {
                    method: "GET",
                    mode: "cors" // no-cors, *cors, same-origin
                }
            ).then(res => res.json());
            set(list);
        },
        create: async (newTodo: Todo) => {
            let result  = await fetch("http://localhost:8080/api/todo",
                {
                    method: "POST",
                    mode: "cors", // no-cors, *cors, same-origin
                    headers: {
                        'Content-Type': 'application/json'
                        // 'Content-Type': 'application/x-www-form-urlencoded',
                    },
                    body: JSON.stringify(newTodo)
                }
            ).then(res => res.json());
            update(current => {
                current.push(result);
                return current;
            });
        },
        modify: async (updated: Todo) => {
            let result  = await fetch(`http://localhost:8080/api/todo/${updated.id}`,
                {
                    method: "PUT",
                    mode: "cors", // no-cors, *cors, same-origin
                    headers: {
                        'Content-Type': 'application/json'
                        // 'Content-Type': 'application/x-www-form-urlencoded',
                    },
                    body: JSON.stringify(updated),
                }
            ).then(res => res.json());
            update(current => {
                current.push(result);
                return current;
            });
        },
        delete: async (id:number) => {
            var status = await fetch(`http://localhost:8080/api/todo/${id}`,
                {
                    method: "DELETE",
                    mode: "cors" // no-cors, *cors, same-origin
                }
            ).then(res => res.status);
            if (status == 200)
            {
                update(current => {
                    return current.filter(v => v.id != id);
                });    
            }
        }
	};
}

export const todoList = createTodoStore();

export class Todo {
    id: number;
    task: string;
    finished: boolean;

    constructor(task: string, id?: number, finished?: boolean) {
        this.task = task;
        this.id = id ?? null;
        this.finished = finished ?? false;
    }
}
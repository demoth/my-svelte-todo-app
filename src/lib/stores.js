import { writable } from 'svelte/store';

// Load todos from localStorage when creating the store
const storedTodos = localStorage.getItem('todos');
const initialTodos = storedTodos ? JSON.parse(storedTodos) : [];

// Create a custom store that syncs with localStorage
function createTodosStore() {
    const { subscribe, set, update } = writable(initialTodos);

    return {
        subscribe,
        set: (value) => {
            localStorage.setItem('todos', JSON.stringify(value));
            set(value);
        },
        update: (updater) => {
            update(todos => {
                const updatedTodos = updater(todos);
                localStorage.setItem('todos', JSON.stringify(updatedTodos));
                return updatedTodos;
            });
        }
    };
}

export const todos = createTodosStore();

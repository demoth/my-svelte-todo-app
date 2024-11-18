import { writable, get } from 'svelte/store';

const API_URL = 'http://localhost:8080/api';

// Create a custom store that syncs with the API
function createTodosStore() {
    const { subscribe, set, update } = writable([]);

    return {
        subscribe,
        set,
        // Load initial todos
        async load() {
            try {
                const response = await fetch(`${API_URL}/todos`);
                const todos = await response.json();
                set(todos);
            } catch (error) {
                console.error('Error loading todos:', error);
            }
        },
        // Add a new todo
        async add(text) {
            try {
                const response = await fetch(`${API_URL}/todos`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({ text, completed: false }),
                });
                const newTodo = await response.json();
                update(todos => [...todos, newTodo]);
            } catch (error) {
                console.error('Error adding todo:', error);
            }
        },
        // Toggle todo completion
        async toggle(id) {
            try {
                const todos = get(this);
                const todo = todos.find(t => t.id === id);
                if (!todo) return;

                const response = await fetch(`${API_URL}/todos/${id}`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        ...todo,
                        completed: !todo.completed,
                        completedAt: !todo.completed ? new Date().toLocaleString() : null,
                    }),
                });
                const updatedTodo = await response.json();
                update(todos => todos.map(t => t.id === id ? updatedTodo : t));
            } catch (error) {
                console.error('Error updating todo:', error);
            }
        },
        // Delete a todo
        async delete(id) {
            try {
                await fetch(`${API_URL}/todos/${id}`, {
                    method: 'DELETE',
                });
                update(todos => todos.filter(t => t.id !== id));
            } catch (error) {
                console.error('Error deleting todo:', error);
            }
        }
    };
}

export const todos = createTodosStore();

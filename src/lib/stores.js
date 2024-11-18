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
                console.log('Loaded todos:', todos);
                set(todos);
            } catch (error) {
                console.error('Error loading todos:', error);
            }
        },
        // Add a new todo
        async add(text) {
            try {
                const todoCreate = {
                    title: text,
                    description: ''  // Optional empty description
                };
                console.log('Sending todo create:', todoCreate);
                
                const response = await fetch(`${API_URL}/todos`, {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(todoCreate),
                });
                const newTodo = await response.json();
                console.log('Added new todo:', newTodo);
                
                if (newTodo.error) {
                    console.error('Error from server:', newTodo.error);
                    return;
                }
                
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

                const todoUpdate = {
                    completed: !todo.completed
                };

                const response = await fetch(`${API_URL}/todos/${id}`, {
                    method: 'PUT',
                    headers: {
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify(todoUpdate),
                });
                const updatedTodo = await response.json();
                console.log('Updated todo:', updatedTodo);
                
                if (updatedTodo.error) {
                    console.error('Error from server:', updatedTodo.error);
                    return;
                }
                
                update(todos => todos.map(t => t.id === id ? updatedTodo : t));
            } catch (error) {
                console.error('Error updating todo:', error);
            }
        },
        // Delete a todo
        async delete(id) {
            try {
                const response = await fetch(`${API_URL}/todos/${id}`, {
                    method: 'DELETE',
                });
                
                if (response.status === 204) {
                    console.log('Deleted todo:', id);
                    update(todos => todos.filter(t => t.id !== id));
                } else {
                    const error = await response.json();
                    console.error('Error from server:', error);
                }
            } catch (error) {
                console.error('Error deleting todo:', error);
            }
        }
    };
}

export const todos = createTodosStore();

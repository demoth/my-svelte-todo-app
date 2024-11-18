<script>
  import { todos } from './stores.js';
  import { onMount } from 'svelte';

  let newTodoText = '';

  function addTodo() {
    if (newTodoText.trim()) {
      todos.add(newTodoText);
      newTodoText = '';
    }
  }

  function deleteTodo(id) {
    todos.delete(id);
  }

  function toggleTodo(id) {
    todos.toggle(id);
  }

  // Load todos when component mounts
  onMount(() => {
    todos.load();
  });
</script>

<div class="todo-list">
  <form on:submit|preventDefault={addTodo}>
    <input 
      type="text" 
      bind:value={newTodoText} 
      placeholder="Enter a new todo" 
    />
    <button type="submit">Add Todo</button>
  </form>

  <div class="todos">
    {#each $todos as todo (todo.id)}
      <div class="todo-item">
        <input
          type="checkbox"
          checked={todo.completed}
          on:change={() => toggleTodo(todo.id)}
        />
        <span class:completed={todo.completed}>{todo.title}</span>
        <button on:click={() => deleteTodo(todo.id)}>Delete</button>
      </div>
    {/each}
  </div>
</div>

<style>
  .todo-list {
    max-width: 500px;
    margin: 0 auto;
  }

  .todo-item {
    display: flex;
    align-items: center;
    gap: 10px;
    padding: 10px;
    border-bottom: 1px solid #eee;
  }

  .completed {
    text-decoration: line-through;
    color: #888;
  }

  form {
    display: flex;
    gap: 10px;
    margin-bottom: 20px;
  }

  input[type="text"] {
    flex: 1;
    padding: 8px;
    border: 1px solid #ddd;
    border-radius: 4px;
  }

  button {
    padding: 8px 16px;
    background-color: #4CAF50;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }

  button:hover {
    background-color: #45a049;
  }

  .todo-item button {
    margin-left: auto;
    background-color: #f44336;
  }

  .todo-item button:hover {
    background-color: #da190b;
  }
</style>

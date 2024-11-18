<script>
  import { todos } from './stores.js';
  import TodoItem from './TodoItem.svelte';
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

  {#each $todos as todo (todo.id)}
    <TodoItem 
      todo={todo} 
      onDelete={() => deleteTodo(todo.id)}
      onToggle={() => toggleTodo(todo.id)}
    />
  {/each}
</div>

<style>
  .todo-list {
    background-color: #f4f4f4;
    border-radius: 8px;
    padding: 20px;
  }

  form {
    display: flex;
    margin-bottom: 20px;
  }

  input {
    flex-grow: 1;
    padding: 10px;
    margin-right: 10px;
    border: 1px solid #ddd;
    border-radius: 4px;
  }

  button {
    padding: 10px 15px;
    background-color: #4CAF50;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
  }

  button:hover {
    background-color: #45a049;
  }
</style>

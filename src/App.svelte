<script>
  import TodoList from './lib/TodoList.svelte';
  import { todos } from './lib/stores.js';
  
  // Calculate statistics
  $: totalItems = $todos.length;
  $: completedItems = $todos.filter(todo => todo.completed).length;
  $: pendingItems = totalItems - completedItems;
</script>

<main>
  <h1>Daniil's Svelte Todo App</h1>
  <TodoList />
</main>

<footer>
  <div class="stats">
    <div class="stat-item">
      <span class="stat-label">Total:</span>
      <span class="stat-value">{totalItems}</span>
    </div>
    <div class="stat-item">
      <span class="stat-label">Completed:</span>
      <span class="stat-value completed">{completedItems}</span>
    </div>
    <div class="stat-item">
      <span class="stat-label">Pending:</span>
      <span class="stat-value pending">{pendingItems}</span>
    </div>
  </div>
</footer>

<style>
  main {
    max-width: 500px;
    margin: 0 auto;
    padding: 20px;
    font-family: Arial, sans-serif;
  }
  h1 {
    text-align: center;
    color: #333;
  }
  footer {
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0;
    background-color: #f4f4f4;
    padding: 15px;
    text-align: center;
    border-top: 1px solid #ddd;
    font-family: Arial, sans-serif;
  }
  .stats {
    display: flex;
    justify-content: center;
    gap: 20px;
  }
  .stat-item {
    display: flex;
    gap: 5px;
  }
  .stat-label {
    color: #666;
    font-weight: 500;
  }
  .stat-value {
    font-weight: bold;
    min-width: 24px;
  }
  .stat-value.completed {
    color: #4CAF50;
  }
  .stat-value.pending {
    color: #f44336;
  }
</style>

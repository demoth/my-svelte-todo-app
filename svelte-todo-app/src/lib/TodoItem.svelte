<script>
  export let todo;
  export let onDelete;
  export let onToggle;
</script>

<div class="todo-item" class:completed={todo.completed}>
  <div class="todo-content">
    <input 
      type="checkbox" 
      checked={todo.completed} 
      on:change={onToggle}
    />
    <span class="todo-text">{todo.text}</span>
    <span class="timestamp">Created: {todo.created}</span>
  </div>
  <button on:click={onDelete}>Delete</button>
</div>

<style>
  .todo-item {
    display: flex;
    align-items: center;
    justify-content: space-between;
    background-color: white;
    padding: 10px;
    margin-bottom: 10px;
    border-radius: 4px;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  }

  .todo-content {
    display: flex;
    align-items: center;
    gap: 10px;
    flex-grow: 1;
  }

  .todo-item.completed .todo-text {
    text-decoration: line-through;
    color: #888;
  }

  input[type="checkbox"] {
    margin-right: 10px;
  }

  .todo-text {
    flex-grow: 1;
  }

  .timestamp {
    font-size: 0.8em;
    color: #666;
    white-space: nowrap;
  }

  button {
    background-color: #f44336;
    color: white;
    border: none;
    padding: 5px 10px;
    border-radius: 4px;
    cursor: pointer;
    margin-left: 10px;
  }

  button:hover {
    background-color: #d32f2f;
  }
</style>

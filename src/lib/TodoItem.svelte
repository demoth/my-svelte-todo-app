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
    <div class="text-content">
      <span class="todo-text">{todo.text}</span>
      <div class="timestamps">
        <span class="timestamp">Created: {todo.created}</span>
        {#if todo.completedAt}
          <span class="timestamp completed">Completed: {todo.completedAt}</span>
        {/if}
      </div>
    </div>
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
    align-items: flex-start;
    gap: 10px;
    flex-grow: 1;
  }

  .text-content {
    display: flex;
    flex-direction: column;
    gap: 4px;
    flex-grow: 1;
  }

  .todo-item.completed .todo-text {
    text-decoration: line-through;
    color: #888;
  }

  input[type="checkbox"] {
    margin-right: 10px;
    margin-top: 4px;
  }

  .todo-text {
    flex-grow: 1;
  }

  .timestamps {
    display: flex;
    flex-direction: column;
    gap: 2px;
  }

  .timestamp {
    font-size: 0.8em;
    color: #666;
    white-space: nowrap;
  }

  .timestamp.completed {
    color: #4CAF50;
  }

  button {
    background-color: #f44336;
    color: white;
    border: none;
    padding: 5px 10px;
    border-radius: 4px;
    cursor: pointer;
    margin-left: 10px;
    align-self: center;
  }

  button:hover {
    background-color: #d32f2f;
  }
</style>

<script>
  export let todo;
  export let onDelete;
  export let onToggle;
</script>

<div class="todo-item" class:completed={todo.completed}>
  <input 
    type="checkbox" 
    checked={todo.completed} 
    on:change={onToggle}
  />
  <span>{todo.text}</span>
  <button on:click={onDelete}>Delete</button>
</div>

<style>
  .todo-item {
    display: flex;
    align-items: center;
    background-color: white;
    padding: 10px;
    margin-bottom: 10px;
    border-radius: 4px;
    box-shadow: 0 2px 4px rgba(0,0,0,0.1);
  }

  .todo-item.completed span {
    text-decoration: line-through;
    color: #888;
  }

  input[type="checkbox"] {
    margin-right: 10px;
  }

  span {
    flex-grow: 1;
  }

  button {
    background-color: #f44336;
    color: white;
    border: none;
    padding: 5px 10px;
    border-radius: 4px;
    cursor: pointer;
  }

  button:hover {
    background-color: #d32f2f;
  }
</style>

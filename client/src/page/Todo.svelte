<script lang="ts">
  import { onMount } from "svelte";
  import { todoList, Todo } from "../store/todo";

  let newTask = "";
  onMount(async () => {
    await todoList.load();
  });
  function add(event)
  {
    var todo = new Todo(newTask);

    todoList.create(todo);
    newTask = "";
  }
</script>

<ul>
  <li><input type="text" bind:value={newTask}/><butoon on:click={add}>add</butoon></li>
  {#each $todoList as todo, i}
	<li>
    <p>
      {todo.task}
    </p>
    <input type=checkbox bind:checked={todo.finished}>
  </li>
{/each}
</ul>

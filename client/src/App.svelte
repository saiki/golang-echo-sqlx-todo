<script lang="ts">
  import Tailwindcss from "./Tailwindcss.svelte";
  import { onMount } from "svelte";

  export let url: string;

  onMount(async () => {
    var result = await fetch(
      "https://dog.ceo/api/breeds/image/random",
      {
        mode: "cors" // no-cors, *cors, same-origin
      }
    ).then((response) => response.json());
    console.log(result);
    url = result.message;
  });
</script>

<Tailwindcss />

<h1>Todo</h1>
<img src={url} />

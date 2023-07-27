<script>
  import svelteLogo from './assets/svelte.svg'
  import viteLogo from '/vite.svg'
  import Counter from './lib/Counter.svelte'

  let ShowVite = false
  let arrayUsers = [
    {
      id: 1,
      framework: "svelte",
      meta: "sveltekit"
    },
    {
      id: 2,
      framework: "react",
      meta: "nextjs"
    },
    {
      id: 3,
      framework: "vue",
      meta: "nuxtjs"
    }
  ]


  const toogle = () => {
    ShowVite = !ShowVite
    arrayUsers = [...arrayUsers, {
      id: arrayUsers.length + 1,
      framework: "starlette",
      meta: "fastapi"
    }]

  }

</script>

<main>
  <div>
    {#if ShowVite}
      <a href="https://vitejs.dev" target="_blank" rel="noreferrer">
      <img src={viteLogo} class="logo" alt="Vite Logo" />
    </a>
      <a href="https://svelte.dev" target="_blank" rel="noreferrer">
      <img src={svelteLogo} class="logo svelte" alt="Svelte Logo" />
    </a>
    {:else}
      <a href="https://svelte.dev" target="_blank" rel="noreferrer">
        <img src={svelteLogo} class="logo svelte" alt="Svelte Logo" />
      </a>
    {/if}


    </div>
    <h1>Vite + Svelte</h1>

    <div class="card">
      <Counter />
      </div>
      <button on:click={toogle}>lalaland</button>

      <p>
        Check out <a href="https://github.com/sveltejs/kit#readme" target="_blank" rel="noreferrer">SvelteKit</a>, the official Svelte app framework powered by Vite!
        </p>

        {#each arrayUsers as au (au.id)}

          <div>{au.id}: {au.framework} meta framework is {au.meta}</div>

        {/each}


        <p class="read-the-docs">
          Click on the Vite and Svelte logos to learn more
  </p>
</main>

<style>
  .logo {
    height: 6em;
    padding: 1.5em;
    will-change: filter;
    transition: filter 300ms;
  }
  .logo:hover {
    filter: drop-shadow(0 0 2em #646cffaa);
  }
  .logo.svelte:hover {
    filter: drop-shadow(0 0 2em #ff3e00aa);
  }
  .read-the-docs {
    color: #888;
  }
</style>

<script>
	import ProjectRow from './ProjectRow.svelte';

	let { projects, selectedProjectId, onselect, oncreate, onrename, ondelete } = $props();

	let newProjectTitle = $state('');
	let creating = $state(false);

	let canCreate = $derived(newProjectTitle.trim().length > 0 && !creating);

	async function create(event) {
		event.preventDefault();

		const title = newProjectTitle.trim();

		if (!title || creating) {
			return;
		}

		creating = true;

		try {
			await oncreate(title);
			newProjectTitle = '';
		} catch {
			// The application displays the error.
		} finally {
			creating = false;
		}
	}
</script>

<aside>
	<h2>Projects</h2>

	<form class="new-project" onsubmit={create}>
		<input
			bind:value={newProjectTitle}
			placeholder="New project"
			aria-label="New project name"
			disabled={creating}
		/>

		<button type="submit" disabled={!canCreate}>
			{creating ? 'Adding…' : 'Add'}
		</button>
	</form>

	<nav aria-label="Projects">
		{#each projects as project (project.id)}
			<ProjectRow
				{project}
				selected={project.id === selectedProjectId}
				{onselect}
				{onrename}
				{ondelete}
			/>
		{/each}
	</nav>

	{#if projects.length === 0}
		<p class="empty">Create your first project above.</p>
	{/if}
</aside>

<style>
	aside {
		min-width: 0;
		padding: 1.25rem;
		color: white;
		background: #111827;
	}

	h2 {
		margin: 0 0 1rem;
	}

	.new-project {
		display: flex;
		gap: 0.5rem;
	}

	input {
		flex: 1;
		min-width: 0;
		padding: 0.7rem;
		color: white;
		border: 1px solid #4b5563;
		border-radius: 7px;
		background: #1f2937;
		font: inherit;
	}

	button {
		padding: 0.7rem;
		color: white;
		border: 1px solid #2563eb;
		border-radius: 7px;
		background: #2563eb;
		font: inherit;
		cursor: pointer;
	}

	button:disabled {
		cursor: not-allowed;
		opacity: 0.5;
	}

	nav {
		display: grid;
		gap: 0.4rem;
		margin-top: 1.25rem;
	}

	.empty {
		color: #9ca3af;
	}
</style>

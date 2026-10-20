<script>
	import ConfirmAction from './ConfirmAction.svelte';
	import InlineEditor from './InlineEditor.svelte';

	let { project, selected = false, onselect, onrename, ondelete } = $props();

	let mode = $state('view');

	async function saveRename(title) {
		if (title === project.title) {
			mode = 'view';
			return;
		}

		await onrename(project.id, title);
		mode = 'view';
	}

	async function confirmDelete() {
		await ondelete(project.id);
		mode = 'view';
	}
</script>

<div class="project-row" class:selected>
	{#if mode === 'rename'}
		<InlineEditor
			initialValue={project.title}
			label="Project name"
			onsave={saveRename}
			oncancel={() => (mode = 'view')}
		/>
	{:else if mode === 'delete'}
		<ConfirmAction
			message={`Delete "${project.title}" and all its todos?`}
			confirmLabel="Delete"
			onconfirm={confirmDelete}
			oncancel={() => (mode = 'view')}
		/>
	{:else}
		<button type="button" class="project-select" onclick={() => onselect(project.id)}>
			<span class="project-title">{project.title}</span>
			<small>{project.todos?.length ?? 0}</small>
		</button>

		<button
			type="button"
			class="icon-button"
			onclick={() => (mode = 'rename')}
			title="Rename project"
			aria-label={`Rename ${project.title}`}
		>
			✎
		</button>

		<button
			type="button"
			class="icon-button danger"
			onclick={() => (mode = 'delete')}
			title="Delete project"
			aria-label={`Delete ${project.title}`}
		>
			×
		</button>
	{/if}
</div>

<style>
	.project-row {
		display: flex;
		align-items: center;
		gap: 0.1rem;
		padding: 0.25rem;
		border-radius: 8px;
	}

	.project-row:hover,
	.project-row.selected {
		background: #374151;
	}

	.project-select {
		display: flex;
		flex: 1;
		align-items: center;
		justify-content: space-between;
		min-width: 0;
		padding: 0.65rem;
		color: inherit;
		border: 0;
		background: transparent;
		text-align: left;
		cursor: pointer;
	}

	.project-title {
		overflow: hidden;
		text-overflow: ellipsis;
		white-space: nowrap;
	}

	small {
		margin-left: 0.5rem;
		color: #9ca3af;
	}

	.icon-button {
		padding: 0.45rem;
		color: inherit;
		border: 0;
		background: transparent;
		cursor: pointer;
	}

	.danger {
		color: #f87171;
	}

	:global(.project-row form input) {
		color: white;
		border-color: #6b7280;
		background: #1f2937;
	}

	:global(.project-row .confirmation) {
		padding: 0.4rem;
	}

	:global(.project-row .confirmation .actions) {
		flex-wrap: wrap;
	}
</style>

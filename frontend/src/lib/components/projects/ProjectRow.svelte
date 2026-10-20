<script>
	import ConfirmAction from '../ui/ConfirmAction.svelte';
	import InlineEditor from '../ui/InlineEditor.svelte';

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

<div
	class={[
		'group flex min-w-0 items-center rounded-xl transition-colors',
		selected
			? 'bg-primary-container text-on-primary-container'
			: 'text-on-surface-variant hover:bg-surface-container-high hover:text-on-surface'
	]}
>
	{#if mode === 'rename'}
		<div class="w-full p-2">
			<InlineEditor
				initialValue={project.title}
				label="Project name"
				onsave={saveRename}
				oncancel={() => (mode = 'view')}
			/>
		</div>
	{:else if mode === 'delete'}
		<div class="w-full p-3">
			<ConfirmAction
				message={`Delete "${project.title}" and its todos?`}
				onconfirm={confirmDelete}
				oncancel={() => (mode = 'view')}
			/>
		</div>
	{:else}
		<button
			type="button"
			class="flex min-w-0 flex-1 items-center justify-between gap-2 px-3 py-2 text-left"
			onclick={() => onselect(project.id)}
		>
			<span class="flex min-w-0 items-center gap-2.5">
				<span class="bg-primary h-2.5 w-2.5 shrink-0 rounded-full"></span>

				<span class="truncate text-sm">
					{project.title}
				</span>
			</span>

			<span class="text-on-surface-variant text-xs">
				{project.todos?.length ?? 0}
			</span>
		</button>

		<div class="mr-1 hidden items-center group-focus-within:flex group-hover:flex">
			<button
				type="button"
				class="hover:bg-surface-container h-7 w-7 rounded-lg"
				title="Rename project"
				aria-label={`Rename ${project.title}`}
				onclick={() => (mode = 'rename')}
			>
				<span class="material-symbols-outlined text-sm"> edit </span>
			</button>

			<button
				type="button"
				class="text-error hover:bg-error-container h-7 w-7 rounded-lg"
				title="Delete project"
				aria-label={`Delete ${project.title}`}
				onclick={() => (mode = 'delete')}
			>
				<span class="material-symbols-outlined text-sm"> delete </span>
			</button>
		</div>
	{/if}
</div>

<script>
	import AppBrand from '../shell/AppBrand.svelte';
	import PrimaryNavigation from '../shell/PrimaryNavigation.svelte';
	import SettingsNavigation from '../shell/SettingsNavigation.svelte';
	import ProjectRow from './ProjectRow.svelte';

	let { projects, selectedProjectId, onselect, oncreate, onrename, ondelete } = $props();

	let creatingProject = $state(false);
	let savingProject = $state(false);
	let newProjectTitle = $state('');
	let projectInput = $state();

	let canCreate = $derived(newProjectTitle.trim().length > 0 && !savingProject);

	$effect(() => {
		if (creatingProject && projectInput) {
			projectInput.focus();
		}
	});

	function openProjectComposer() {
		creatingProject = true;
	}

	function closeProjectComposer() {
		creatingProject = false;
		newProjectTitle = '';
	}

	async function submitProject(event) {
		event.preventDefault();

		const title = newProjectTitle.trim();

		if (!title || savingProject) {
			return;
		}

		savingProject = true;

		try {
			await oncreate(title);
			closeProjectComposer();
		} catch {
			// The app-level error banner handles the error.
		} finally {
			savingProject = false;
		}
	}
</script>

<aside
	class="bg-surface-container-low fixed top-0 left-0 z-50 hidden h-full w-72 flex-col justify-between py-6 shadow-[0_1px_8px_rgba(46,50,48,0.04)] lg:flex"
>
	<div class="flex min-h-0 flex-1 flex-col space-y-6 px-4">
		<AppBrand />

		<PrimaryNavigation />

		<section class="flex min-h-0 flex-1 flex-col pt-4">
			<div class="mb-2 flex items-center justify-between px-3">
				<span class="text-secondary text-xs font-semibold tracking-wider uppercase">
					Projects
				</span>

				<button
					class="text-on-surface-variant hover:text-on-surface hover:bg-surface-container-high rounded-md p-1 transition-colors"
					type="button"
					aria-label="Create project"
					onclick={openProjectComposer}
				>
					<span class="material-symbols-outlined text-base"> add </span>
				</button>
			</div>

			{#if creatingProject}
				<form class="bg-surface-container mb-2 rounded-xl p-2" onsubmit={submitProject}>
					<input
						bind:this={projectInput}
						bind:value={newProjectTitle}
						class="bg-surface-container-lowest text-on-surface focus:ring-primary/30 w-full rounded-lg border-0 px-3 py-2 text-sm focus:ring-2 focus:outline-none"
						placeholder="Project name"
						aria-label="Project name"
						disabled={savingProject}
					/>

					<div class="mt-2 flex gap-2">
						<button
							type="submit"
							class="bg-primary text-on-primary rounded-lg px-3 py-1.5 text-xs disabled:opacity-50"
							disabled={!canCreate}
						>
							{savingProject ? 'Creating…' : 'Create'}
						</button>

						<button
							type="button"
							class="text-on-surface-variant hover:bg-surface-container-high rounded-lg px-3 py-1.5 text-xs"
							disabled={savingProject}
							onclick={closeProjectComposer}
						>
							Cancel
						</button>
					</div>
				</form>
			{/if}

			<nav class="min-h-0 space-y-1 overflow-y-auto" aria-label="Projects">
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

			{#if projects.length === 0 && !creatingProject}
				<p class="text-on-surface-variant px-3 py-2 text-xs">Create your first project to begin.</p>
			{/if}
		</section>
	</div>

	<div class="px-4 pt-4">
		<SettingsNavigation />
	</div>
</aside>

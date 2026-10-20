<script>
	import {
		createProject as createProjectRequest,
		deleteProject as deleteProjectRequest,
		getProjects,
		renameProject as renameProjectRequest
	} from '../lib/api/projects.js';

	import {
		createTodo as createTodoRequest,
		deleteTodo as deleteTodoRequest,
		updateTodo
	} from '../lib/api/todos.js';

	import ErrorBanner from '../lib/components/ErrorBanner.svelte';
	import ProjectSidebar from '../lib/components/ProjectSidebar.svelte';
	import TodoList from '../lib/components/TodoList.svelte';

	/** @type {import("../lib/domain/types.js").Project[]} */
	let projects = $state([]);

	let selectedProjectId = $state(null);
	let error = $state('');
	let loading = $state(true);

	let selectedProject = $derived(
		projects.find((project) => project.id === selectedProjectId) ?? null
	);

	$effect(() => {
		refreshProjects(null).catch(() => {
			// The error is already shown in the page.
		});
	});

	function getErrorMessage(value) {
		return value instanceof Error ? value.message : 'An unexpected error occurred';
	}

	async function execute(action) {
		error = '';

		try {
			return await action();
		} catch (requestError) {
			error = getErrorMessage(requestError);
			throw requestError;
		}
	}

	async function refreshProjects(preferredProjectId = selectedProjectId) {
		return execute(async () => {
			const loadedProjects = await getProjects();

			projects = loadedProjects;

			const preferredProjectExists = loadedProjects.some(
				(project) => project.id === preferredProjectId
			);

			selectedProjectId = preferredProjectExists
				? preferredProjectId
				: (loadedProjects[0]?.id ?? null);

			loading = false;
		});
	}

	function selectProject(projectId) {
		selectedProjectId = projectId;
	}

	async function createProject(title) {
		return execute(async () => {
			const project = await createProjectRequest(title);
			await refreshProjects(project.id);
		});
	}

	async function renameProject(projectId, title) {
		return execute(async () => {
			await renameProjectRequest(projectId, title);
			await refreshProjects(projectId);
		});
	}

	async function deleteProject(projectId) {
		return execute(async () => {
			await deleteProjectRequest(projectId);
			await refreshProjects();
		});
	}

	async function createTodo(title) {
		if (!selectedProject) {
			return;
		}

		return execute(async () => {
			await createTodoRequest(selectedProject.id, title);
			await refreshProjects(selectedProject.id);
		});
	}

	async function toggleTodo(todoId, done) {
		return execute(async () => {
			await updateTodo(todoId, { done });
			await refreshProjects(selectedProjectId);
		});
	}

	async function renameTodo(todoId, title) {
		return execute(async () => {
			await updateTodo(todoId, { title });
			await refreshProjects(selectedProjectId);
		});
	}

	async function deleteTodo(todoId) {
		return execute(async () => {
			await deleteTodoRequest(todoId);
			await refreshProjects(selectedProjectId);
		});
	}
</script>

<svelte:head>
	<title>Quick Todo</title>
	<meta name="description" content="A simple project-based todo application" />
</svelte:head>

<div class="app">
	<ProjectSidebar
		{projects}
		{selectedProjectId}
		onselect={selectProject}
		oncreate={createProject}
		onrename={renameProject}
		ondelete={deleteProject}
	/>

	<main>
		<ErrorBanner message={error} ondismiss={() => (error = '')} />

		{#if loading}
			<div class="status">
				<p>Loading projects…</p>
			</div>
		{:else if selectedProject}
			<header>
				<div>
					<p class="eyebrow">Project</p>
					<h1>{selectedProject.title}</h1>
				</div>
			</header>

			<TodoList
				todos={selectedProject.todos ?? []}
				oncreate={createTodo}
				ontoggle={toggleTodo}
				onrename={renameTodo}
				ondelete={deleteTodo}
			/>
		{:else}
			<div class="status">
				<h1>No projects yet</h1>
				<p>Create your first project in the sidebar.</p>
			</div>
		{/if}
	</main>
</div>

<style>
	:global(*) {
		box-sizing: border-box;
	}

	:global(body) {
		margin: 0;
		color: #1f2937;
		background: #f3f4f6;
		font-family: system-ui, sans-serif;
	}

	:global(button),
	:global(input) {
		font: inherit;
	}

	.app {
		display: grid;
		grid-template-columns: 300px minmax(0, 1fr);
		min-height: 100vh;
	}

	main {
		width: 100%;
		max-width: 950px;
		padding: 3rem;
	}

	header {
		margin-bottom: 2rem;
	}

	h1 {
		margin: 0;
	}

	.eyebrow {
		margin: 0 0 0.25rem;
		color: #6b7280;
		font-size: 0.8rem;
		font-weight: bold;
		letter-spacing: 0.08em;
		text-transform: uppercase;
	}

	.status {
		margin-top: 20vh;
		color: #6b7280;
		text-align: center;
	}

	.status h1 {
		color: #1f2937;
	}

	@media (max-width: 750px) {
		.app {
			grid-template-columns: 1fr;
		}

		main {
			padding: 1.5rem;
		}
	}
</style>

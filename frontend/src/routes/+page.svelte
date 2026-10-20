<script>
	import DashboardHero from '../lib/components/dashboard/DashboardHero.svelte';
	import InspirationBanner from '../lib/components/dashboard/InspirationBanner.svelte';
	import TaskDetailPanel from '../lib/components/dashboard/TaskDetailPanel.svelte';
	import ProjectSidebar from '../lib/components/projects/ProjectSidebar.svelte';
	import TopBar from '../lib/components/shell/TopBar.svelte';
	import TodoList from '../lib/components/todos/TodoList.svelte';
	import ErrorBanner from '../lib/components/ui/ErrorBanner.svelte';
	import { TodoAppState } from '../lib/state/todo-app.svelte.js';

	const app = new TodoAppState();

	let selectedTodos = $derived(app.selectedProject?.todos ?? []);

	let selectedCompletedCount = $derived(selectedTodos.filter((todo) => todo.done).length);

	let selectedCompletionPercentage = $derived(
		selectedTodos.length === 0
			? 0
			: Math.round((selectedCompletedCount / selectedTodos.length) * 100)
	);

	$effect(() => {
		void app.load();
	});
</script>

<svelte:head>
	<title>Terra Tasks</title>

	<meta name="description" content="A calm, project-based todo application" />
</svelte:head>

<ProjectSidebar
	projects={app.projects}
	selectedProjectId={app.selectedProjectId}
	onselect={(projectId) => app.selectProject(projectId)}
	oncreate={(title) => app.addProject(title)}
	onrename={(projectId, title) => app.editProject(projectId, title)}
	ondelete={(projectId) => app.removeProject(projectId)}
/>

<div class="lg:pl-72">
	<TopBar />

	<main class="bg-surface relative min-h-screen w-full pt-16">
		<div class="mx-auto w-full max-w-7xl px-4 py-8 sm:px-8">
			<ErrorBanner message={app.error} ondismiss={() => app.dismissError()} />

			{#if app.loading}
				<div class="flex min-h-[60vh] flex-col items-center justify-center gap-3">
					<div
						class="border-primary h-10 w-10 animate-spin rounded-full border-4 border-t-transparent"
					></div>

					<p class="text-on-surface-variant text-sm">Preparing your workspace…</p>
				</div>
			{:else if app.selectedProject}
				<DashboardHero
					project={app.selectedProject}
					completed={selectedCompletedCount}
					total={selectedTodos.length}
					percentage={selectedCompletionPercentage}
				/>

				<div class="grid grid-cols-1 items-start gap-8 xl:grid-cols-12">
					<div class="flex flex-col gap-8 xl:col-span-7">
						<TodoList
							todos={selectedTodos}
							projectTitle={app.selectedProject.title}
							focusToken={app.composerFocusToken}
							oncreate={(title) => app.addTodo(title)}
							ontoggle={(todoId, done) => app.toggleTodo(todoId, done)}
							onrename={(todoId, title) => app.editTodo(todoId, title)}
							ondelete={(todoId) => app.removeTodo(todoId)}
						/>

						<InspirationBanner />
					</div>

					<TaskDetailPanel project={app.selectedProject} />
				</div>
			{:else}
				<div class="flex min-h-[65vh] flex-col items-center justify-center text-center">
					<div
						class="bg-primary-fixed text-primary mb-5 flex h-16 w-16 items-center justify-center rounded-2xl"
					>
						<span class="material-symbols-outlined text-3xl"> psychiatry </span>
					</div>

					<h1 class="font-headline text-on-surface text-3xl font-semibold">
						Cultivate your first project
					</h1>

					<p class="text-on-surface-variant mt-2 max-w-md text-sm">
						Use the plus button beside Projects in the sidebar to create a place for your tasks.
					</p>
				</div>
			{/if}
		</div>
	</main>
</div>

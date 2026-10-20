<!DOCTYPE html>

<html lang="en"
	><head
		><meta charset="utf-8" /><meta
			content="width=device-width, initial-scale=1.0"
			name="viewport"
		/><meta content="web_dashboard" name="shell-type" /><link
			href="https://fonts.googleapis.com"
			rel="preconnect"
		/><link crossorigin="" href="https://fonts.gstatic.com" rel="preconnect" /><link
			href="https://fonts.googleapis.com/css2?family=Literata:ital,opsz,wght@0,7..72,400;0,7..72,600;0,7..72,700;1,7..72,400&amp;family=Nunito+Sans:ital,wght@0,400;0,600;0,700;1,400&amp;display=swap"
			rel="stylesheet"
		/><link
			href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:opsz,wght,FILL,GRAD@20..48,100..700,0..1,-50..200"
			rel="stylesheet"
		/>
		<link
			href="https://fonts.googleapis.com/css2?family=Material+Symbols+Outlined:wght,FILL@100..700,0..1&amp;display=swap"
			rel="stylesheet"
		/><style>
			@layer base {
				html,
				body {
					margin: 0;
					padding: 0;
				}
				body {
					overscroll-behavior: none;
				}
				main > :first-child {
					margin-top: 0 !important;
				}
				main > :last-child {
					margin-bottom: 0 !important;
				}
			}
			::-webkit-scrollbar {
				display: none;
			}
		</style><script src="https://cdn.tailwindcss.com"></script><script id="tailwind-config">
			tailwind.config = {
				darkMode: 'class',
				theme: {
					extend: {
						colors: {
							'surface-bright': '#faf6f0',
							'surface-container-high': '#eae6de',
							'primary-fixed-dim': '#8ecf9e',
							'on-primary-fixed-variant': '#2a6038',
							surface: '#faf6f0',
							'tertiary-fixed-dim': '#dcc48e',
							'on-tertiary-fixed': '#221a05',
							'on-surface-variant': '#4a4e4a',
							'tertiary-container': '#c4a66a',
							'on-tertiary-fixed-variant': '#554020',
							'inverse-on-surface': '#f5f0e8',
							'surface-container-lowest': '#ffffff',
							'on-primary': '#ffffff',
							'on-background': '#2e3230',
							'on-tertiary-container': '#554020',
							'surface-container': '#f0ece4',
							outline: '#74796e',
							'surface-tint': '#4a7c59',
							'surface-variant': '#e4e0d8',
							'surface-container-low': '#f5f1ea',
							'on-secondary-fixed-variant': '#4a4538',
							'error-container': '#ffdad8',
							'on-tertiary': '#ffffff',
							primary: '#4a7c59',
							'on-secondary-fixed': '#1e1a13',
							'surface-dim': '#dbd7cf',
							'on-surface': '#2e3230',
							'on-primary-container': '#d8f0de',
							'secondary-fixed': '#f0e8db',
							'on-secondary': '#ffffff',
							'primary-fixed': '#c8e8d0',
							'on-primary-fixed': '#002110',
							'surface-container-highest': '#e4e0d8',
							tertiary: '#705c30',
							'inverse-surface': '#2e3230',
							'on-error-container': '#690005',
							'primary-container': '#78a886',
							'on-secondary-container': '#5e5548',
							'tertiary-fixed': '#f8e0a8',
							'secondary-container': '#f0e8db',
							'secondary-fixed-dim': '#d4ccbf',
							'inverse-primary': '#8ecf9e',
							secondary: '#6b6358',
							'on-error': '#ffffff',
							error: '#b83230',
							background: '#faf6f0',
							'outline-variant': '#c4c8bc'
						},
						borderRadius: { DEFAULT: '0.5rem', lg: '1rem', xl: '1.5rem', full: '9999px' },
						spacing: {},
						fontFamily: {
							headline: ['Literata'],
							display: ['Literata'],
							body: ['Nunito Sans'],
							label: ['Nunito Sans']
						},
						fontSize: {}
					}
				}
			};
		</script></head
	><body class="bg-surface font-body text-on-surface antialiased"
		><aside
			class="bg-surface-container-low fixed top-0 left-0 z-50 flex h-full w-72 flex-col justify-between py-6 shadow-[0_1px_8px_rgba(46,50,48,0.04)]"
		>
			<div class="flex flex-col space-y-6 px-4">
				<div class="flex items-center gap-3 px-3">
					<img
						alt="Terra Tasks Logo"
						class="h-8 w-auto object-contain"
						src="https://lh3.googleusercontent.com/aida-public/AB6AXuBXNhiuEoozLTqclV1SIOV88v17hiQDgWhn5iiWj6-3fPtU5fsF-yIikSSvRw9v1uj6n1y3LkaERi4x_-3_qD-CJgP4JTBb_KNwTnqiPD9fJudImma4Wevfy8K7qEZSD-ULACYUaWaoiOI1j9ycOvYzAYoRFg8kAcWlonQ89dYUq9ew3x1Ibz5zw5QQUN-pcm0FrMA1i1uopC4hhqOm2zgO41dQK3cbfqfxmQQ3pW2E5GkBFRJbgt8"
					/><span class="font-headline text-on-surface text-xl font-semibold tracking-tight"
						>Terra Tasks</span
					>
				</div>
				<div class="px-2">
					<button
						class="bg-primary text-on-primary hover:bg-primary/90 flex w-full items-center justify-center gap-2 rounded-xl px-4 py-2.5 font-medium shadow-[0_4px_20px_rgba(46,50,48,0.06)] transition-colors"
						type="button"
						><span class="material-symbols-outlined text-lg">add_circle</span><span>Add Task</span
						></button
					>
				</div>
				<nav
					class="space-y-1"
					data-active-classes="bg-primary-container text-on-primary-container font-semibold rounded-xl"
				>
					<a
						aria-current="page"
						class="bg-primary-container text-on-primary-container flex items-center justify-between rounded-xl px-3 py-2.5 font-semibold transition-colors"
						data-path="today"
						href="#"
						><div class="flex items-center gap-3">
							<span class="material-symbols-outlined text-xl">wb_sunny</span><span class="text-sm"
								>Today</span
							>
						</div>
						<span
							class="bg-surface-container text-on-surface-variant rounded-full px-2 py-0.5 text-xs font-semibold"
							>4</span
						></a
					><a
						class="text-on-surface-variant hover:bg-surface-container-high hover:text-on-surface flex items-center justify-between rounded-xl px-3 py-2.5 transition-colors"
						data-path="inbox"
						href="#"
						><div class="flex items-center gap-3">
							<span class="material-symbols-outlined text-xl">inbox</span><span class="text-sm"
								>Inbox</span
							>
						</div>
						<span
							class="bg-surface-container text-on-surface-variant rounded-full px-2 py-0.5 text-xs font-semibold"
							>12</span
						></a
					><a
						class="text-on-surface-variant hover:bg-surface-container-high hover:text-on-surface flex items-center justify-between rounded-xl px-3 py-2.5 transition-colors"
						data-path="upcoming"
						href="#"
						><div class="flex items-center gap-3">
							<span class="material-symbols-outlined text-xl">calendar_month</span><span
								class="text-sm">Upcoming</span
							>
						</div></a
					><a
						class="text-on-surface-variant hover:bg-surface-container-high hover:text-on-surface flex items-center justify-between rounded-xl px-3 py-2.5 transition-colors"
						data-path="completed"
						href="#"
						><div class="flex items-center gap-3">
							<span class="material-symbols-outlined text-xl">check_circle</span><span
								class="text-sm">Completed</span
							>
						</div></a
					>
				</nav>
				<div class="pt-4">
					<div class="mb-2 flex items-center justify-between px-3">
						<span class="text-secondary text-xs font-semibold tracking-wider uppercase"
							>Projects</span
						><button
							class="text-on-surface-variant hover:text-on-surface hover:bg-surface-container-high rounded-md p-1 transition-colors"
							type="button"><span class="material-symbols-outlined text-base">add</span></button
						>
					</div>
					<nav
						class="space-y-1"
						data-active-classes="bg-primary-container text-on-primary-container font-semibold rounded-xl"
					>
						<a
							class="text-on-surface-variant hover:bg-surface-container-high hover:text-on-surface flex items-center justify-between rounded-xl px-3 py-2 transition-colors"
							data-path="project-gardening"
							href="#"
							><div class="flex items-center gap-2.5">
								<span class="bg-primary h-2.5 w-2.5 rounded-full"></span><span
									class="truncate text-sm">Greenhouse Flora</span
								>
							</div>
							<span class="text-on-surface-variant text-xs">6</span></a
						><a
							class="text-on-surface-variant hover:bg-surface-container-high hover:text-on-surface flex items-center justify-between rounded-xl px-3 py-2 transition-colors"
							data-path="project-homestead"
							href="#"
							><div class="flex items-center gap-2.5">
								<span class="bg-tertiary h-2.5 w-2.5 rounded-full"></span><span
									class="truncate text-sm">Homestead Build</span
								>
							</div>
							<span class="text-on-surface-variant text-xs">14</span></a
						><a
							class="text-on-surface-variant hover:bg-surface-container-high hover:text-on-surface flex items-center justify-between rounded-xl px-3 py-2 transition-colors"
							data-path="project-autumn-harvest"
							href="#"
							><div class="flex items-center gap-2.5">
								<span class="bg-tertiary-container h-2.5 w-2.5 rounded-full"></span><span
									class="truncate text-sm">Autumn Harvest</span
								>
							</div>
							<span class="text-on-surface-variant text-xs">3</span></a
						>
					</nav>
				</div>
			</div>
			<div class="px-4 pt-4">
				<nav
					class="space-y-1"
					data-active-classes="bg-primary-container text-on-primary-container font-semibold rounded-xl"
				>
					<a
						class="text-on-surface-variant hover:bg-surface-container-high hover:text-on-surface flex items-center gap-3 rounded-xl px-3 py-2.5 transition-colors"
						data-path="settings"
						href="#"
						><span class="material-symbols-outlined text-xl">settings</span><span class="text-sm"
							>Settings</span
						></a
					>
				</nav>
			</div>
		</aside>
		<div class="pl-72">
			<header
				class="bg-surface/80 fixed top-0 right-0 left-72 z-40 flex h-16 items-center justify-between px-8 shadow-[0_1px_8px_rgba(46,50,48,0.04)] backdrop-blur-xl"
			>
				<div class="flex w-96 items-center">
					<div class="relative flex w-full items-center">
						<span
							class="material-symbols-outlined text-on-surface-variant pointer-events-none absolute left-3 text-lg"
							>search</span
						><input
							class="bg-surface-container-low text-on-surface placeholder-on-surface-variant focus:ring-primary/40 w-full rounded-xl border-0 py-2 pr-4 pl-9 text-sm transition-all focus:ring-2 focus:outline-none"
							placeholder="Search tasks, projects, tags..."
							type="text"
						/>
					</div>
				</div>
				<div class="flex items-center gap-4">
					<button
						class="text-on-surface-variant hover:bg-surface-container-high hover:text-on-surface flex h-9 w-9 items-center justify-center rounded-xl transition-colors"
						type="button"
						><span class="material-symbols-outlined text-xl">notifications</span></button
					><button
						class="text-on-surface-variant hover:bg-surface-container-high hover:text-on-surface flex h-9 w-9 items-center justify-center rounded-xl transition-colors"
						type="button"><span class="material-symbols-outlined text-xl">tune</span></button
					>
					<div class="flex items-center gap-3 pl-2">
						<img
							alt="Profile"
							class="h-8 w-8 rounded-full object-cover"
							src="https://lh3.googleusercontent.com/aida-public/AB6AXuASVVMTTj5KO_lFuWyz4vjXWo69ujOSbkDelVnw9ZF3KsTSq2nJZW1vP3_BywawpoVlXVdXD3t2LlDwn1KhkOhCcbj5--VBFfKgYwel50ycKpndgvnWBZztDzP-7HUnCnarBRwi4gsWTLOdJJklgt9OwhV-3JuuJke8Pb8VU9QpJXo2y9toCXdJ3iOV2NiUxg5OqLRvFzxWTuXqIczU7ajyQa0CDTtbNuu8nUqv5z-e5pyPROloEN8"
						/>
						<div class="hidden flex-col text-left sm:flex">
							<span class="text-on-surface text-xs leading-tight font-semibold">Elena Rostova</span
							><span class="text-on-surface-variant text-[11px] leading-tight">Pro Plan</span>
						</div>
					</div>
				</div>
			</header>
			<main class="bg-surface relative min-h-screen w-full pt-16">
				<div class="flex w-full flex-col">
					<div class="mx-auto w-full max-w-7xl px-8 py-8">
						<!-- Top Greeting & Dynamic Metrics -->
						<div class="mb-8 flex flex-col justify-between gap-6 lg:flex-row lg:items-end">
							<div>
								<div class="mb-2 flex items-center gap-2.5">
									<span
										class="bg-primary-fixed text-on-primary-fixed inline-flex items-center gap-1.5 rounded-full px-3 py-1 text-xs font-semibold tracking-wide"
									>
										<span class="material-symbols-outlined text-sm">filter_vintage</span>
										Daily Rhythm
									</span>
									<span
										class="text-secondary font-label text-xs font-medium tracking-wider uppercase"
										>Thursday, October 24</span
									>
								</div>
								<h1
									class="font-headline text-on-surface text-3xl font-semibold tracking-tight sm:text-4xl"
								>
									Today's Focus
								</h1>
								<p class="text-on-surface-variant font-body mt-1.5 text-sm">
									Ground yourself in deliberate, thoughtful progress.
								</p>
							</div>
							<!-- Warm motivational progress card -->
							<div class="bg-surface-container flex items-center gap-5 rounded-xl p-4 shadow-sm">
								<div class="relative flex h-14 w-14 shrink-0 items-center justify-center">
									<svg class="h-14 w-14 -rotate-90 transform" viewbox="0 0 48 48">
										<circle
											class="text-surface-container-high"
											cx="24"
											cy="24"
											fill="transparent"
											r="20"
											stroke="currentColor"
											stroke-width="4"
										></circle>
										<circle
											class="text-primary transition-all duration-700"
											cx="24"
											cy="24"
											fill="transparent"
											r="20"
											stroke="currentColor"
											stroke-dasharray="125.6"
											stroke-dashoffset="47.7"
											stroke-linecap="round"
											stroke-width="4.5"
										></circle>
									</svg>
									<span class="text-on-surface font-body absolute text-xs font-semibold">62%</span>
								</div>
								<div>
									<div class="flex items-center gap-2">
										<span class="text-on-surface font-headline text-sm font-semibold"
											>5 of 8 completed</span
										>
										<span
											class="material-symbols-outlined text-primary text-base"
											style="font-variation-settings: 'FILL' 1;">spa</span
										>
									</div>
									<p class="text-on-surface-variant font-body mt-0.5 text-xs">
										Steady pace. You're past the peak for today.
									</p>
								</div>
							</div>
						</div>
						<!-- Quick Add Inline Field -->
						<div class="mb-8">
							<div
								class="bg-surface-container-lowest relative flex items-center gap-3 rounded-xl p-2 shadow-sm transition-all focus-within:shadow-md"
							>
								<div
									class="bg-primary-fixed text-primary ml-1 flex h-9 w-9 shrink-0 items-center justify-center rounded-full"
								>
									<span class="material-symbols-outlined text-xl">add</span>
								</div>
								<input
									class="text-on-surface placeholder:text-outline font-body w-full bg-transparent py-1.5 text-sm focus:outline-none"
									id="quickTaskInput"
									placeholder="Add a mindful task for today... (press Enter to harvest)"
									type="text"
								/>
								<div class="flex items-center gap-2 pr-2">
									<button
										class="bg-surface-container text-on-surface-variant hover:bg-surface-container-high hidden items-center gap-1 rounded-lg px-2.5 py-1 text-xs transition-colors sm:inline-flex"
										type="button"
									>
										<span class="material-symbols-outlined text-sm">schedule</span>
										<span>Today 2:00 PM</span>
									</button>
									<button
										class="bg-surface-container text-on-surface-variant hover:bg-surface-container-high hidden items-center gap-1 rounded-lg px-2.5 py-1 text-xs transition-colors sm:inline-flex"
										type="button"
									>
										<span class="material-symbols-outlined text-tertiary text-sm">bookmark</span>
										<span>Gardening</span>
									</button>
									<button
										class="bg-primary text-on-primary hover:bg-primary/90 rounded-lg px-3.5 py-2 text-xs font-medium transition-colors"
										type="button"
									>
										Cultivate
									</button>
								</div>
							</div>
						</div>
						<!-- Main Content Area: Grouped Task Lists + Slide-Over Task Detail Panel -->
						<div class="grid grid-cols-1 items-start gap-8 xl:grid-cols-12">
							<!-- Left: Grouped Task Stream (7 cols on xl, or full if detail panel closed) -->
							<div class="flex flex-col gap-8 xl:col-span-7">
								<!-- SECTION 1: Morning Focus -->
								<section class="space-y-3">
									<div class="flex items-center justify-between pb-1">
										<div class="flex items-center gap-2">
											<span class="material-symbols-outlined text-tertiary text-lg">light_mode</span
											>
											<h2
												class="font-headline text-on-surface text-base font-semibold tracking-tight"
											>
												Morning Focus
											</h2>
										</div>
										<span
											class="bg-surface-container text-secondary rounded-full px-2 py-0.5 text-xs font-semibold"
											>2 / 2 done</span
										>
									</div>
									<div class="space-y-2">
										<!-- Completed Task Card -->
										<div
											class="group bg-surface-container-low flex items-center justify-between gap-4 rounded-xl p-4 opacity-75 shadow-sm transition-all hover:opacity-100"
										>
											<div class="flex min-w-0 items-center gap-3.5">
												<button
													class="bg-primary text-on-primary flex h-5 w-5 shrink-0 items-center justify-center rounded-full transition-transform active:scale-90"
													type="button"
												>
													<span class="material-symbols-outlined text-sm font-bold">check</span>
												</button>
												<div class="min-w-0">
													<span
														class="text-outline font-body block truncate text-sm font-medium line-through"
														>Inspect soil moisture &amp; root aeration in Sector 4</span
													>
													<div class="mt-1 flex items-center gap-2">
														<span
															class="bg-primary-fixed text-on-primary-fixed-variant rounded-full px-2 py-0.5 text-[11px] font-semibold"
															>Greenhouse Flora</span
														>
														<span class="text-outline flex items-center gap-1 text-[11px]">
															<span class="material-symbols-outlined text-xs">done_all</span> 8:30 AM
														</span>
													</div>
												</div>
											</div>
											<span class="material-symbols-outlined text-outline/50 text-base"
												>drag_indicator</span
											>
										</div>
										<!-- Completed Task Card 2 -->
										<div
											class="group bg-surface-container-low flex items-center justify-between gap-4 rounded-xl p-4 opacity-75 shadow-sm transition-all hover:opacity-100"
										>
											<div class="flex min-w-0 items-center gap-3.5">
												<button
													class="bg-primary text-on-primary flex h-5 w-5 shrink-0 items-center justify-center rounded-full transition-transform active:scale-90"
													type="button"
												>
													<span class="material-symbols-outlined text-sm font-bold">check</span>
												</button>
												<div class="min-w-0">
													<span
														class="text-outline font-body block truncate text-sm font-medium line-through"
														>Refill rainwater drip irrigation manifolds</span
													>
													<div class="mt-1 flex items-center gap-2">
														<span
															class="bg-tertiary-fixed text-on-tertiary-fixed-variant rounded-full px-2 py-0.5 text-[11px] font-semibold"
															>Homestead Build</span
														>
														<span class="text-outline flex items-center gap-1 text-[11px]">
															<span class="material-symbols-outlined text-xs">done_all</span> 9:45 AM
														</span>
													</div>
												</div>
											</div>
											<span class="material-symbols-outlined text-outline/50 text-base"
												>drag_indicator</span
											>
										</div>
									</div>
								</section>
								<!-- SECTION 2: Afternoon Work -->
								<section class="space-y-3">
									<div class="flex items-center justify-between pb-1">
										<div class="flex items-center gap-2">
											<span class="material-symbols-outlined text-primary text-lg">wb_sunny</span>
											<h2
												class="font-headline text-on-surface text-base font-semibold tracking-tight"
											>
												Afternoon Work
											</h2>
										</div>
										<span
											class="bg-surface-container text-on-surface-variant rounded-full px-2 py-0.5 text-xs font-semibold"
											>2 remaining</span
										>
									</div>
									<div class="space-y-2.5">
										<!-- Active Selected Task Card -->
										<div
											class="bg-surface-container-lowest relative flex cursor-pointer items-start justify-between gap-4 overflow-hidden rounded-xl p-4 shadow-md transition-all duration-200"
										>
											<div
												class="bg-primary absolute top-0 bottom-0 left-0 w-1.5 rounded-l-xl"
											></div>
											<div class="flex min-w-0 items-start gap-3.5 pl-1">
												<button
													class="border-primary/60 hover:border-primary mt-0.5 flex h-5 w-5 shrink-0 items-center justify-center rounded-full border-2 transition-colors"
													type="button"
												>
													<!-- Custom unchecked circle -->
												</button>
												<div class="min-w-0">
													<div class="flex items-center gap-2">
														<span
															class="text-on-surface font-body block truncate text-sm font-semibold"
															>Harvest heirloom heirloom squash &amp; categorize seeds</span
														>
														<span
															class="material-symbols-outlined text-error shrink-0 text-base"
															style="font-variation-settings: 'FILL' 1;"
															title="High Priority">flag</span
														>
													</div>
													<p class="text-on-surface-variant font-body mt-1 line-clamp-1 text-xs">
														Check heirloom bin tags against seed vault ledger before sun drops.
													</p>
													<div class="mt-2.5 flex flex-wrap items-center gap-2">
														<span
															class="bg-tertiary-fixed text-on-tertiary-fixed-variant rounded-full px-2 py-0.5 text-[11px] font-semibold"
															>Autumn Harvest</span
														>
														<span
															class="bg-surface-container text-on-surface-variant flex items-center gap-1 rounded-full px-2 py-0.5 text-[11px] font-medium"
														>
															<span class="material-symbols-outlined text-xs">schedule</span> 2:30 PM
														</span>
														<span
															class="bg-surface-container text-on-surface-variant flex items-center gap-1 rounded-full px-2 py-0.5 text-[11px] font-medium"
														>
															<span class="material-symbols-outlined text-xs">checklist</span> 3 subtasks
														</span>
													</div>
												</div>
											</div>
											<span class="material-symbols-outlined text-primary mt-0.5 shrink-0 text-xl"
												>chevron_right</span
											>
										</div>
										<!-- Task Card 4 -->
										<div
											class="group bg-surface-container-low hover:bg-surface-container-lowest flex cursor-pointer items-start justify-between gap-4 rounded-xl p-4 shadow-sm transition-all"
										>
											<div class="flex min-w-0 items-start gap-3.5">
												<button
													class="border-outline hover:border-primary mt-0.5 flex h-5 w-5 shrink-0 items-center justify-center rounded-full border-2 transition-colors"
													type="button"
												></button>
												<div class="min-w-0">
													<div class="flex items-center gap-2">
														<span
															class="text-on-surface font-body block truncate text-sm font-medium"
															>Install secondary solar inverter in potting shed</span
														>
														<span
															class="material-symbols-outlined text-tertiary shrink-0 text-base"
															style="font-variation-settings: 'FILL' 1;"
															title="Medium Priority">flag</span
														>
													</div>
													<div class="mt-2 flex flex-wrap items-center gap-2">
														<span
															class="bg-secondary-fixed text-on-secondary-fixed-variant rounded-full px-2 py-0.5 text-[11px] font-semibold"
															>Homestead Build</span
														>
														<span
															class="bg-surface-container text-on-surface-variant flex items-center gap-1 rounded-full px-2 py-0.5 text-[11px] font-medium"
														>
															<span class="material-symbols-outlined text-xs">schedule</span> 4:15 PM
														</span>
														<span
															class="bg-surface-container text-on-surface-variant flex items-center gap-1 rounded-full px-2 py-0.5 text-[11px] font-medium"
														>
															<span class="material-symbols-outlined text-xs">checklist</span> 4 subtasks
														</span>
													</div>
												</div>
											</div>
											<span
												class="material-symbols-outlined text-outline/40 group-hover:text-on-surface-variant text-base"
												>more_vert</span
											>
										</div>
									</div>
								</section>
								<!-- SECTION 3: Personal & Home -->
								<section class="space-y-3">
									<div class="flex items-center justify-between pb-1">
										<div class="flex items-center gap-2">
											<span class="material-symbols-outlined text-secondary text-lg">cottage</span>
											<h2
												class="font-headline text-on-surface text-base font-semibold tracking-tight"
											>
												Personal &amp; Home
											</h2>
										</div>
										<span
											class="bg-surface-container text-on-surface-variant rounded-full px-2 py-0.5 text-xs font-semibold"
											>1 remaining</span
										>
									</div>
									<div class="space-y-2.5">
										<!-- Task Card 5 -->
										<div
											class="group bg-surface-container-low hover:bg-surface-container-lowest flex cursor-pointer items-start justify-between gap-4 rounded-xl p-4 shadow-sm transition-all"
										>
											<div class="flex min-w-0 items-start gap-3.5">
												<button
													class="border-outline hover:border-primary mt-0.5 flex h-5 w-5 shrink-0 items-center justify-center rounded-full border-2 transition-colors"
													type="button"
												></button>
												<div class="min-w-0">
													<div class="flex items-center gap-2">
														<span
															class="text-on-surface font-body block truncate text-sm font-medium"
															>Steep lavender &amp; chamomile botanical tea blend</span
														>
													</div>
													<div class="mt-2 flex flex-wrap items-center gap-2">
														<span
															class="bg-primary-fixed text-on-primary-fixed-variant rounded-full px-2 py-0.5 text-[11px] font-semibold"
															>Botanicals</span
														>
														<span
															class="bg-surface-container text-on-surface-variant flex items-center gap-1 rounded-full px-2 py-0.5 text-[11px] font-medium"
														>
															<span class="material-symbols-outlined text-xs">schedule</span> 7:00 PM
														</span>
													</div>
												</div>
											</div>
											<span
												class="material-symbols-outlined text-outline/40 group-hover:text-on-surface-variant text-base"
												>more_vert</span
											>
										</div>
										<!-- Task Card 6 Completed -->
										<div
											class="group bg-surface-container-low flex items-center justify-between gap-4 rounded-xl p-4 opacity-75 shadow-sm transition-all hover:opacity-100"
										>
											<div class="flex min-w-0 items-center gap-3.5">
												<button
													class="bg-primary text-on-primary flex h-5 w-5 shrink-0 items-center justify-center rounded-full transition-transform active:scale-90"
													type="button"
												>
													<span class="material-symbols-outlined text-sm font-bold">check</span>
												</button>
												<div class="min-w-0">
													<span
														class="text-outline font-body block truncate text-sm font-medium line-through"
														>Journal 3 seasonal micro-observations</span
													>
													<div class="mt-1 flex items-center gap-2">
														<span
															class="bg-secondary-fixed text-on-secondary-fixed-variant rounded-full px-2 py-0.5 text-[11px] font-semibold"
															>Mindfulness</span
														>
														<span class="text-outline flex items-center gap-1 text-[11px]">
															<span class="material-symbols-outlined text-xs">done_all</span> 7:15 AM
														</span>
													</div>
												</div>
											</div>
											<span class="material-symbols-outlined text-outline/50 text-base"
												>drag_indicator</span
											>
										</div>
									</div>
								</section>
								<!-- Ambient Inspiration Banner -->
								<div
									class="from-primary-fixed/40 via-surface-container to-surface-container-high flex items-center gap-4 rounded-xl bg-gradient-to-r p-5"
								>
									<div
										class="bg-surface-container-lowest text-primary flex h-12 w-12 shrink-0 items-center justify-center rounded-xl shadow-sm"
									>
										<span class="material-symbols-outlined text-2xl">nest_eco_leaf</span>
									</div>
									<div>
										<span
											class="text-primary font-label text-xs font-semibold tracking-wider uppercase"
											>Homestead Thought</span
										>
										<p
											class="text-on-surface-variant font-body mt-0.5 text-xs leading-relaxed italic"
										>
											"Nature does not hurry, yet everything is accomplished." Take a deep breath
											between chores.
										</p>
									</div>
								</div>
							</div>
							<!-- Right: Detailed Slide-Over / Inspect Panel (5 cols on xl) -->
							<aside
								class="bg-surface-container-lowest sticky top-24 flex flex-col gap-6 rounded-xl p-6 shadow-sm xl:col-span-5"
							>
								<!-- Panel Header -->
								<div class="flex items-center justify-between border-b-0 pb-4">
									<div class="flex items-center gap-2">
										<span class="bg-tertiary h-2.5 w-2.5 rounded-full"></span>
										<span
											class="text-secondary font-label text-xs font-semibold tracking-wider uppercase"
											>Autumn Harvest</span
										>
									</div>
									<div class="flex items-center gap-1.5">
										<button
											class="text-on-surface-variant hover:bg-surface-container flex h-8 w-8 items-center justify-center rounded-lg transition-colors"
											title="Mark favorite"
											type="button"
										>
											<span class="material-symbols-outlined text-lg">star</span>
										</button>
										<button
											class="text-on-surface-variant hover:bg-surface-container flex h-8 w-8 items-center justify-center rounded-lg transition-colors"
											title="Options"
											type="button"
										>
											<span class="material-symbols-outlined text-lg">more_horiz</span>
										</button>
										<button
											class="text-on-surface-variant hover:bg-surface-container flex h-8 w-8 items-center justify-center rounded-lg transition-colors"
											title="Close Panel"
											type="button"
										>
											<span class="material-symbols-outlined text-lg">close</span>
										</button>
									</div>
								</div>
								<!-- Task Title & Completed Toggle in Panel -->
								<div class="flex items-start gap-3">
									<button
										class="border-primary/70 hover:border-primary mt-1 flex h-6 w-6 shrink-0 items-center justify-center rounded-full border-2 transition-colors"
										type="button"
									>
									</button>
									<div class="flex-1">
										<h3 class="font-headline text-on-surface text-lg leading-snug font-semibold">
											Harvest heirloom heirloom squash &amp; categorize seeds
										</h3>
										<p class="text-on-surface-variant font-body mt-1 text-xs">
											Created at 6:40 AM by Elena Rostova
										</p>
									</div>
								</div>
								<!-- Metadata Attributes Grid -->
								<div class="bg-surface-container-low grid grid-cols-2 gap-3 rounded-xl p-3.5">
									<!-- Due Date Picker Pill -->
									<div class="flex flex-col gap-1">
										<span
											class="text-secondary font-label text-[11px] font-semibold tracking-wider uppercase"
											>Due Date</span
										>
										<button
											class="bg-surface-container-lowest text-on-surface hover:bg-surface-container flex items-center gap-2 rounded-lg px-3 py-2 text-left text-xs font-medium transition-colors"
											type="button"
										>
											<span class="material-symbols-outlined text-primary text-sm"
												>calendar_today</span
											>
											<span>Today, 2:30 PM</span>
										</button>
									</div>
									<!-- Priority Selector Pill -->
									<div class="flex flex-col gap-1">
										<span
											class="text-secondary font-label text-[11px] font-semibold tracking-wider uppercase"
											>Priority</span
										>
										<button
											class="bg-surface-container-lowest text-error hover:bg-surface-container flex items-center gap-2 rounded-lg px-3 py-2 text-left text-xs font-semibold transition-colors"
											type="button"
										>
											<span
												class="material-symbols-outlined text-sm"
												style="font-variation-settings: 'FILL' 1;">flag</span
											>
											<span>High Priority</span>
										</button>
									</div>
									<!-- Project / Category -->
									<div class="flex flex-col gap-1">
										<span
											class="text-secondary font-label text-[11px] font-semibold tracking-wider uppercase"
											>Project</span
										>
										<button
											class="bg-surface-container-lowest text-on-surface hover:bg-surface-container flex items-center gap-2 rounded-lg px-3 py-2 text-left text-xs font-medium transition-colors"
											type="button"
										>
											<span class="bg-tertiary h-2 w-2 rounded-full"></span>
											<span class="truncate">Autumn Harvest</span>
										</button>
									</div>
									<!-- Tags -->
									<div class="flex flex-col gap-1">
										<span
											class="text-secondary font-label text-[11px] font-semibold tracking-wider uppercase"
											>Tags</span
										>
										<div class="flex items-center gap-1">
											<span
												class="bg-surface-container-lowest text-on-surface rounded-md px-2 py-1 text-[11px] font-medium"
												>#seed-bank</span
											>
											<button
												class="bg-surface-container-lowest hover:bg-surface-container text-outline flex h-6 w-6 items-center justify-center rounded-md text-xs"
												type="button"
											>
												+
											</button>
										</div>
									</div>
								</div>
								<!-- Notes Section -->
								<div class="space-y-2">
									<div class="flex items-center justify-between">
										<span
											class="text-secondary font-label text-xs font-semibold tracking-wider uppercase"
											>Field Notes</span
										>
										<span class="text-outline text-[11px] font-medium">Auto-saved</span>
									</div>
									<div
										class="bg-surface-container-low focus-within:bg-surface-container rounded-xl p-3.5 transition-colors"
									>
										<textarea
											class="text-on-surface placeholder:text-outline font-body w-full resize-none bg-transparent text-xs leading-relaxed focus:outline-none"
											placeholder="Add observations, soil notes, or reminders..."
											rows="3"
											>Store seeds in glass mason jars with silica packets. Note down row numbers for blue hubbard squash vs buttercup varieties.</textarea
										>
									</div>
								</div>
								<!-- Subtasks Checklist -->
								<div class="space-y-3">
									<div class="flex items-center justify-between">
										<span
											class="text-secondary font-label text-xs font-semibold tracking-wider uppercase"
											>Subtasks (1 of 3)</span
										>
										<button
											class="text-primary flex items-center gap-0.5 text-xs font-semibold hover:underline"
											type="button"
										>
											<span class="material-symbols-outlined text-sm">add</span> New
										</button>
									</div>
									<div class="space-y-2">
										<!-- Completed Subtask -->
										<div class="bg-surface-container-low flex items-center gap-3 rounded-lg p-2">
											<button
												class="bg-primary text-on-primary flex h-4 w-4 shrink-0 items-center justify-center rounded-full"
												type="button"
											>
												<span class="material-symbols-outlined text-[10px] font-bold">check</span>
											</button>
											<span class="text-outline font-body flex-1 text-xs line-through"
												>Sanitize garden shears with herbal wash</span
											>
										</div>
										<!-- Pending Subtask 1 -->
										<div
											class="hover:bg-surface-container-low group flex items-center gap-3 rounded-lg p-2 transition-colors"
										>
											<button
												class="border-outline hover:border-primary h-4 w-4 shrink-0 rounded-full border"
												type="button"
											></button>
											<span class="text-on-surface font-body flex-1 text-xs"
												>Cut stems leaving 2 inches for dry curing</span
											>
											<button
												class="text-outline hover:text-error opacity-0 transition-opacity group-hover:opacity-100"
												type="button"
											>
												<span class="material-symbols-outlined text-xs">delete</span>
											</button>
										</div>
										<!-- Pending Subtask 2 -->
										<div
											class="hover:bg-surface-container-low group flex items-center gap-3 rounded-lg p-2 transition-colors"
										>
											<button
												class="border-outline hover:border-primary h-4 w-4 shrink-0 rounded-full border"
												type="button"
											></button>
											<span class="text-on-surface font-body flex-1 text-xs"
												>Log final weight in cellar ledger</span
											>
											<button
												class="text-outline hover:text-error opacity-0 transition-opacity group-hover:opacity-100"
												type="button"
											>
												<span class="material-symbols-outlined text-xs">delete</span>
											</button>
										</div>
									</div>
								</div>
								<!-- Visual Attachment Preview -->
								<div class="space-y-2">
									<span
										class="text-secondary font-label text-xs font-semibold tracking-wider uppercase"
										>Visual Reference</span
									>
									<div
										class="bg-surface-container relative h-32 overflow-hidden rounded-xl shadow-sm"
									>
										<img
											class="h-full w-full object-cover"
											data-alt="A warm top-down organic photograph of freshly harvested winter squash, orange pumpkins, and rustic gardening shears on a weathered wooden potting bench, surrounded by dried herbs under gentle afternoon sunlight with earthy tones."
											src="https://lh3.googleusercontent.com/aida-public/AB6AXuAamAdFgt1AdWLB8lP-CVBj4VrjWwsUImO2_oHy-60QlyPqOzBopgplCVadFvUHg6sZUesu-tNIAUK9iiNlcDVaX10DzAyPv4-fBAOEJRM4X-prl62Yd7Kq3HPc8JlWLZLB0U_P_H9XU-ZfdBztWjSi21CX7LPTziDAj2K28jpasRZ7QnMC0PvSERStRuHt64cGFFptkEP2W2jWeBpsmc2ozSFw_lbVQl4VJgety8lnL2X0zby7wYI"
										/>
										<div
											class="from-inverse-surface/60 absolute inset-0 flex items-end bg-gradient-to-t via-transparent to-transparent p-2.5"
										>
											<span class="text-inverse-on-surface truncate text-[11px] font-medium"
												>Harvest_heirloom_cellar_spec.jpg</span
											>
										</div>
									</div>
								</div>
								<!-- Activity History -->
								<div class="space-y-2 border-t-0 pt-3">
									<span
										class="text-secondary font-label text-xs font-semibold tracking-wider uppercase"
										>Activity</span
									>
									<div class="text-on-surface-variant font-body space-y-2 text-[11px]">
										<div class="flex items-center gap-2">
											<span class="bg-primary h-1.5 w-1.5 rounded-full"></span>
											<span>Elena completed <strong>Sanitize shears</strong> at 11:20 AM</span>
										</div>
										<div class="flex items-center gap-2">
											<span class="bg-tertiary h-1.5 w-1.5 rounded-full"></span>
											<span>Priority shifted to <strong>High</strong> at 10:15 AM</span>
										</div>
									</div>
								</div>
								<!-- Panel Footer Actions -->
								<div class="flex items-center justify-between pt-2">
									<button
										class="text-error flex items-center gap-1 text-xs font-medium hover:underline"
										type="button"
									>
										<span class="material-symbols-outlined text-sm">delete_outline</span>
										<span>Archive</span>
									</button>
									<button
										class="bg-primary text-on-primary hover:bg-primary/90 rounded-xl px-4 py-2 text-xs font-medium shadow-sm transition-colors"
										type="button"
									>
										Save Changes
									</button>
								</div>
							</aside>
						</div>
					</div>
				</div>
			</main>
		</div></body
	></html
>

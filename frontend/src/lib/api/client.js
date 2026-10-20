import { env } from '$env/dynamic/public';

export class ApiError extends Error {
	constructor(message, status) {
		super(message);
		this.name = 'ApiError';
		this.status = status;
	}
}

function getApiBaseUrl() {
	return env.PUBLIC_API_URL?.replace(/\/$/, '') || 'http://localhost:8080/api';
}

async function getErrorMessage(response) {
	const contentType = response.headers.get('Content-Type') ?? '';

	if (contentType.includes('application/json')) {
		try {
			const body = await response.json();
			return body.error || body.message || 'Request failed';
		} catch {
			return 'Request failed';
		}
	}

	return (await response.text()) || 'Request failed';
}

export async function apiRequest(path, options = {}) {
	const response = await fetch(`${getApiBaseUrl()}${path}`, {
		...options,
		headers: {
			...(options.body ? { 'Content-Type': 'application/json' } : {}),
			...options.headers
		}
	});

	if (!response.ok) {
		throw new ApiError(await getErrorMessage(response), response.status);
	}

	if (response.status === 204) {
		return null;
	}

	return response.json();
}

const API_BASE_URL = import.meta.env.VITE_API_URL ?? 'http://localhost:8080/api';

export class ApiError extends Error {
	constructor(message, status) {
		super(message);
		this.name = 'ApiError';
		this.status = status;
	}
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
	const response = await fetch(`${API_BASE_URL}${path}`, {
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

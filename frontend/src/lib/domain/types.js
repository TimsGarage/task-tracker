/**
 * @typedef {Object} Subtask
 * @property {number} id
 * @property {number} todoId
 * @property {string} title
 * @property {boolean} done
 */

/**
 * @typedef {Object} Todo
 * @property {number} id
 * @property {number} projectId
 * @property {string} title
 * @property {boolean} done
 * @property {string | null} dueDate
 * @property {"low" | "medium" | "high"} priority
 * @property {string} notes
 * @property {string[]} tags
 * @property {Subtask[]} subtasks
 */

/**
 * @typedef {Object} Project
 * @property {number} id
 * @property {string} title
 * @property {Todo[]} todos
 */

export {};

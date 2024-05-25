import { writable, type Writable } from 'svelte/store';
import type { Organisation } from '../../types/Organisation';
import storage from './storage';

export const currOrganisationStore = storage<Organisation | undefined>(
  'currOrganisation',
  undefined,
);

export const toast_store: Writable<string> = writable('');

export const setToast = (message: string) => {
  toast_store.set('');
  toast_store.set(message);
};

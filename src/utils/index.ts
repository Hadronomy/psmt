export function merge<TElement>(first: TElement, second: Object) : TElement {
  return Object.assign(first, Object.assign(second));
}

export function sleep(ms: number): Promise<void> {
  // eslint-disable-next-line no-promise-executor-return
  return new Promise<void>((resolve) => setTimeout(resolve, ms));
}

export default {
  merge,
};

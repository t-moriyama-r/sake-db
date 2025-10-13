export class AssertionError extends Error {}

export const assertNonNullable: <T>(
  value: T,
) => asserts value is NonNullable<T> = (value) => {
  if (value == null) {
    throw new AssertionError(
      '想定していないエラーです。[value]が存在しません。',
    );
  }
};

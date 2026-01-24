import { describe, expect, it } from 'vitest';

import { AssertionError, assertNonNullable } from './assertNonNullable';

describe('assertNonNullable', () => {
  it('nullの場合にAssertionErrorをスローする', () => {
    expect(() => assertNonNullable(null)).toThrow(AssertionError);
    expect(() => assertNonNullable(null)).toThrow(
      '想定していないエラーです。[value]が存在しません。',
    );
  });

  it('undefinedの場合にAssertionErrorをスローする', () => {
    expect(() => assertNonNullable(undefined)).toThrow(AssertionError);
    expect(() => assertNonNullable(undefined)).toThrow(
      '想定していないエラーです。[value]が存在しません。',
    );
  });

  it('null以外の値の場合は例外をスローしない', () => {
    expect(() => assertNonNullable(0)).not.toThrow();
    expect(() => assertNonNullable('')).not.toThrow();
    expect(() => assertNonNullable(false)).not.toThrow();
    expect(() => assertNonNullable([])).not.toThrow();
    expect(() => assertNonNullable({})).not.toThrow();
    expect(() => assertNonNullable('test')).not.toThrow();
    expect(() => assertNonNullable(123)).not.toThrow();
    expect(() => assertNonNullable(true)).not.toThrow();
  });

  it('型推論が正しく機能する', () => {
    // このテストはコンパイル時の型チェックを確認するためのものです
    const value: string | null = 'test';
    assertNonNullable(value);
    // この時点でvalueの型はstringになっているはず（nullが除外される）
    const result: string = value; // エラーが出なければ型推論が正しく機能している
    expect(result).toBe('test');
  });
});

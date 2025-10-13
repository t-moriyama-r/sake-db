// 型定義として不要な__typenameを消す
export function stripTypeName<
  T extends { __typename: string } = { __typename: string },
>(arg: T): Omit<T, '__typename'> {
  const { __typename: _, ...rest } = arg;
  return rest;
}

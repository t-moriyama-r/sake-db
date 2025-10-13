import type { QueryOptions } from '@apollo/client/core/watchQueryOptions';

import useQuery from '@/funcs/composable/useQuery/useQuery';
import { stripTypeName } from '@/funcs/util/stripTypeName';
import type { CategoryQuery } from '@/graphQL/auto-generated';
import {
  type Category,
  type CategoryResponse,
  GET_DETAIL,
} from '@/graphQL/Category/categories';

export const getCategoryDetail = () => {
  const { fetch } =
    useQuery<CategoryResponse<CategoryQuery['category']>>(GET_DETAIL);
  const fetchCategoryDetail = async (
    request: { id: number },
    options?: Omit<QueryOptions, 'query' | 'variables'>,
  ) => {
    const response = await fetch(request, options);
    const data = stripTypeName(response.category);
    return {
      category: {
        ...data,
        description: data.description ?? '',
      },
    } as { category: Category };
  };

  return { fetch: fetchCategoryDetail };
};

import { type DocumentNode, gql } from '@apollo/client/core';

import { type Liquor as CardLiquor } from '../Index/random';

export interface RelatedLiquorsResponse {
  readonly relatedLiquors: CardLiquor[];
}

export const RELATED_LIQUORS_GET: DocumentNode = gql`
  query relatedLiquors($id: String!) {
    relatedLiquors(liquorId: $id) {
      id
      name
      categoryId
      categoryName
      description
      imageBase64
      updatedAt
    }
  }
`;

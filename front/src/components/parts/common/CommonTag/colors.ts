export const colorClassMap = {
  default: {
    base: 'text-blue-600 border-blue-600',
    hover: 'hover:text-blue-800 hover:border-blue-800',
  },
  primary: {
    base: 'text-blue-600 border-blue-600 bg-white',
    hover: 'hover:text-blue-800 hover:border-blue-800',
  },
  secondary: {
    base: 'text-green-600 border-green-600 bg-white',
    hover: 'hover:text-green-800 hover:border-green-800',
  },
  danger: {
    base: 'text-red-600 border-red-600 bg-white',
    hover: 'hover:text-red-800 hover:border-red-800',
  },
} as const;

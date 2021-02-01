import React from 'react';

const SvgDrop = props => (
  <svg width={props.width || 64} height={props.height || 64} {...props}>
    <path d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16-7.163 16-16 16zm7.134-14.533c-.551-2.427-2.264-5.103-2.264-5.103 1.974 5.387 1.214 9.45.691 11.755 0 0-.69 2.578.795.43 1.852-2.688.778-7.082.778-7.082zM18.258 7.91c1.027 3.658 1.492 6.763 1.637 9.352.255 4.412-.418 7.315-.836 8.848-.134.499.424.87.749.493l.006-.006c.812-.929 1.196-3.105 1.376-5.015l.052-.627c.168-2.3.07-4.621-.47-6.845-.993-4.104-2.514-6.2-2.514-6.2zm-1.84 6.444c-.25-6.008-.703-11.343-.703-11.343s-.458 5.068-.708 10.942c-.04.953-.076 1.922-.105 2.897-.04 1.417-.058 2.84-.058 4.215.006 2.444.1 4.737.32 6.566.087.725 1.027.725 1.114 0 .151-1.248.238-2.781.273-4.465a123.84 123.84 0 0 0-.023-5.462c-.03-1.115-.064-2.241-.11-3.35zm-7.019 7.24a18.405 18.405 0 0 1-.093-3.414c.058-.894.186-1.852.407-2.874.203-.935.481-1.916.859-2.949 0 0-1.208 1.887-1.916 3.924-.14.39-.255.79-.343 1.179 0 0-.226 1.022-.26 2.386-.024.946.051 2.066.353 3.123a6.19 6.19 0 0 0 .68 1.573s1.486 2.212.795-.43a19.786 19.786 0 0 1-.482-2.519zm2.119-1.185a31.681 31.681 0 0 1 .058-3.588 35.673 35.673 0 0 1 .279-2.635c.26-1.881.685-3.965 1.33-6.276 0 0-1.528 2.096-2.52 6.2-.053.22-.1.441-.146.656-.18.894-.296 1.8-.354 2.717-.07 1.15-.052 2.31.03 3.466l.052.633c.18 1.91.557 4.086 1.375 5.015l.006.006c.325.372.888 0 .75-.493-.303-1.132-.75-3.013-.86-5.701zm2.537.226a99.317 99.317 0 0 1-.018-3.994c.012-.876.035-1.8.076-2.769.099-2.502.569-6.119.894-9.184 0 0-1.602 3.315-2.334 9.295-.098.818-.191 1.7-.261 2.653l-.052.708s-.087 1.265-.053 2.972c.047 2.194.29 5.12 1.196 7.036.03.058.064.11.11.15l.006.007c.343.325.871.017.825-.488a81.99 81.99 0 0 1-.39-6.386zm3.32 0c-.075 3.112-.272 5.289-.388 6.386-.047.505.481.813.824.488l.006-.006a.483.483 0 0 0 .11-.151c.906-1.916 1.15-4.842 1.196-7.036.035-1.707-.052-2.972-.052-2.972l-.053-.708a56.374 56.374 0 0 0-.26-2.653c-.732-5.98-2.335-9.295-2.335-9.295.326 3.065.796 6.682.894 9.184.041.97.064 1.893.076 2.77a99.317 99.317 0 0 1-.018 3.993z" />
  </svg>
);

export default SvgDrop;
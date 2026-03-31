# 02. Principal Component Analysis (PCA)

## A. Definition and Purpose

- **Dimensionality Reduction:** A technique used to reduce the number of variables in a dataset while preserving as much information (variance) as possible.
- **Unsupervised Nature:** PCA does not consider target labels; it focuses purely on the feature structure.
- **Key Objectives:**
    1. **Data Compression:** Reduces storage and computational costs.
    2. **Visualization:** Projects high-dimensional data into 2D or 3D spaces.
    3. **Feature Engineering:** Removes multicollinearity by creating uncorrelated features.
    4. **Noise Filtering:** Discards components with very low variance, which often represent noise.

## B. Core Mathematical Concepts

- **Variance:** Measures the spread of data along a single axis.
- **Covariance:** Measures how two variables change together.
- **Principal Components (PCs):**
    - **PC1:** The direction that captures the maximum variance in the data.
    - **PC2:** The direction orthogonal (perpendicular) to PC1 that captures the next highest variance.
- **Orthogonality:** All principal components are uncorrelated with each other.

## C. Step-by-Step Algorithm

### 1. Standardization (Scaling)

- **Requirement:** Variables must be centered and scaled (mean = 0, variance = 1) using Z-score normalization: $z = \frac{x - \mu}{\sigma}$.
- **Reasoning:** Without scaling, a variable with a large range (e.g., Salary) will dominate variables with small ranges (e.g., Age).

### 2. Covariance Matrix Computation

- **Goal:** To identify the relationships and correlations between all pairs of variables in the dataset.
- **Output:** A $d \times d$ matrix (where $d$ is the number of dimensions).

### 3. Eigendecomposition

- **Eigenvectors:** Determine the directions of the new feature space (the axes).
- **Eigenvalues:** Determine the magnitude (amount of variance) explained by each eigenvector.

### 4. Selecting Principal Components

- **Scree Plot:** A graphical tool used to visualize eigenvalues and decide how many components to keep.
- **Explained Variance Ratio:** Calculated as $\frac{\lambda_i}{\sum \lambda}$, representing the percentage of total variance carried by the $i^{th}$ component.

### 5. Transformation (Recasting)

- **Projection:** The original data is multiplied by the selected eigenvectors (the Feature Vector) to transform the data into the new reduced-dimension subspace.

## D. Limitations of PCA

- **Linearity:** PCA assumes linear relationships between variables; it may fail to capture complex non-linear structures (Kernel PCA is used for non-linear cases).
- **Interpretability:** The new principal components are linear combinations of original features, making it hard to explain what a "component" physically represents.
- **Outliers:** PCA is sensitive to outliers as they significantly affect the variance and covariance calculations.

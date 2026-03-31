# 01. Clustering

## A. Introduction to Clustering

- **I. Definition:** An unsupervised learning technique used to group data points based on similarity.
- **II. Goal:** Achieve high **intra-cluster similarity** (points within a group are similar) and low **inter-cluster similarity** (different groups are distinct).
- **III. Unsupervised Nature:** Unlike classification, clustering works on unlabeled data to discover hidden structures.

## B. Core Concepts and Metrics

- **I. Distance Metrics:** The measure of similarity $d(x, y)$ between points.
    - **1. Euclidean Distance:** $d(x, y) = \sqrt{\sum_{i=1}^{n} (x_i - y_i)^2}$
    - **2. Manhattan Distance:** $d(x, y) = \sum_{i=1}^{n} |x_i - y_i|$
    - **3. Cosine Similarity:** $\cos(\theta) = \frac{\mathbf{A} \cdot \mathbf{B}}{\|\mathbf{A}\| \|\mathbf{B}\|}$

- **II. Hard vs. Soft Clustering:**
    - **1. Hard Clustering:** Each point belongs to exactly one cluster (e.g., K-Means).
    - **2. Soft (Fuzzy) Clustering:** Points have a probability of belonging to multiple clusters (e.g., GMM).

## C. Major Clustering Algorithms

### 1. Partitioning Methods

- **a. K-Means:**
    - **i. Process:** Randomly initializes $K$ centroids $\mu_k$, assigns points to the nearest centroid by minimizing $\sum \|x_i - \mu_k\|^2$, and re-calculates centroids until convergence.
    - **ii. Limitations:** Requires $K$ to be specified; sensitive to outliers and initialization.

- **b. K-Medoids (PAM):**
    - **i. Process:** Similar to K-Means but uses actual data points as centers (medoids).
    - **ii. Advantage:** More robust to noise and outliers than K-Means.

### 2. Hierarchical Methods

- **a. Agglomerative (Bottom-Up):** Starts with each point as its own cluster and merges the closest pairs until one cluster remains.
- **b. Divisive (Top-Down):** Starts with one giant cluster and recursively splits it into smaller ones.
- **c. Linkage Criteria:** Defines how distance between clusters is calculated (Single, Complete, or Average Linkage).

### 3. Density-Based Methods

- **a. DBSCAN:**
    - **i. Process:** Groups points that are close to many neighbors and marks points in low-density regions as outliers.
    - **ii. Strengths:** Can find clusters of arbitrary shapes and naturally handles noise.

### 4. Distribution-Based Methods

- **a. Gaussian Mixture Models (GMM):**
    - **i. Process:** Assumes data is a mixture of several Gaussian distributions.
    - **ii. Optimization:** Uses the Expectation-Maximization (EM) algorithm to estimate parameters.

## D. Evaluation and Selection

| Metric               | Description                                                                                    | Formula / Range                                                       |
| :------------------- | :--------------------------------------------------------------------------------------------- | :-------------------------------------------------------------------- |
| **Elbow Method**     | Finds optimal $K$ by identifying the point where the rate of decrease in WCSS sharply changes. | $WCSS = \sum_{i \in C_k} \|x_i - \mu_k\|^2$                           |
| **Silhouette Score** | Measures how well a point fits its cluster vs. the next closest one.                           | $s(i) = \frac{b(i) - a(i)}{\max\{a(i), b(i)\}}$ <br> Range: $[-1, 1]$ |

> **Note:** A Silhouette Score near +1 indicates perfect clustering, while the "Elbow" represents a balance between complexity and error.

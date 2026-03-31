# 04. Recommender Systems

## A. Introduction to Recommender Systems

- **Definition:** Information filtering systems that predict the "rating" or "preference" a user would give to an item.
- **Purpose:** To help users discover relevant content (movies, products, music) and increase platform engagement.
- **Components:**
    - **Users:** Individuals with specific preferences.
    - **Items:** Products, movies, or content being recommended.
    - **Interactions:** Ratings, clicks, views, or purchases.

## B. Major Types of Recommendation Engines

## 1. Content-Based Filtering

- **Mechanism:** Recommends items similar to those a user liked in the past based on **item features**.
- **Process:**
    - **Item Profile:** Keywords, genres, or descriptions (e.g., "Action", "Directed by Nolan").
    - **User Profile:** Aggregated preferences based on interacted items.
- **Pros:** No "Cold Start" problem for new items; highly personalized to a specific user.
- **Cons:** **Filter Bubbles** (limited discovery of new genres); requires manual tagging/metadata.

## 2. Collaborative Filtering (CF)

- **Mechanism:** Based on the assumption that "if User A and User B agree on one issue, they are likely to agree on others."
- **User-Based CF:** Finds users similar to the target user and recommends items they liked.
- **Item-Based CF:** Finds items similar to those the user has liked based on how _other_ users rated them.
- **Matrix Factorization:**
    - **Concept:** Decomposes the large User-Item Interaction Matrix into two lower-dimensional matrices (latent factors).
    - **Algorithms:** Singular Value Decomposition (SVD), Alternating Least Squares (ALS).

## 3. Hybrid Systems

- **Mechanism:** Combines Content-Based and Collaborative Filtering to leverage the strengths of both.
- **Example:** Netflix uses Collaborative Filtering (what similar users watch) combined with Content-Based data (genre, actors).

## C. Evaluation Metrics

- **Precision and Recall:** Measures the accuracy of the recommended list.
- **Root Mean Squared Error (RMSE):** Measures the difference between predicted ratings and actual ratings.
- **Mean Average Precision (MAP):** Evaluates the ranking quality of recommendations.
- **Diversity and Serendipity:** Measures how varied and surprisingly relevant the recommendations are.

## D. Key Challenges

- **Cold Start Problem:** Difficult to recommend for new users or new items with zero interaction history.
- **Data Sparsity:** Most users only interact with a tiny fraction of the total items, making the interaction matrix mostly empty.
- **Scalability:** Calculating similarities for millions of users and items in real-time requires massive computation.

## E. Modern Approaches

- **Deep Learning:** Using Neural Networks (like RNNs or Transformers) to capture sequential patterns in user behavior.
- **Knowledge-Based:** Using explicit rules (e.g., "Suggest a laptop with 16GB RAM for a developer").
- **Reinforcement Learning:** Treating recommendations as a sequence of actions where the system learns to maximize long-term user satisfaction.

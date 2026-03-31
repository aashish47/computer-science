# 03. Anomaly Detection

## A. Introduction to Anomaly Detection

- **Definition:** The process of identifying data points, events, or observations that deviate significantly from the dataset’s normal behavior.
- **Synonyms:** Also referred to as Outlier Detection, Fraud Detection, or Novelty Detection.
- **Key Challenges:**
    - **Imbalanced Data:** Anomalies are rare by definition, making supervised learning difficult.
    - **Evolving Patterns:** What is "normal" can change over time (concept drift).
    - **Boundary Definition:** The line between normal and anomalous is often fuzzy.

## B. Types of Anomalies

- **Point Anomalies:** A single instance that is anomalous relative to the rest of the data (e.g., a single transaction of $50,000 in a history of $20 purchases).
- **Contextual (Conditional) Anomalies:** An instance that is anomalous in a specific context but normal otherwise (e.g., 35°C is normal in summer but an anomaly in winter).
- **Collective Anomalies:** A collection of related data points that are anomalous together, even if individual points are not (e.g., a sequence of small, rapid login attempts indicating a bot attack).

## C. Detection Techniques and Algorithms

### Statistical Methods

- **Z-Score (Standard Score):** Measures how many standard deviations a point is from the mean.
    - **Threshold:** Points with $|Z| > 3$ are typically flagged.

- **Interquartile Range (IQR):**
    - **Formula:** Outliers are points below $Q1 - 1.5 \times IQR$ or above $Q3 + 1.5 \times IQR$.
    - **Benefit:** More robust to existing outliers than the Z-score.

### Proximity-Based Methods

- **K-Nearest Neighbors (KNN):** Points with the largest distance to their $k^{th}$ nearest neighbor are considered anomalies.
- **Local Outlier Factor (LOF):**
    - **Mechanism:** Compares the local density of a point to the local densities of its neighbors.
    - **Strength:** Can find outliers in datasets with varying densities.

### Tree-Based Methods

- **Isolation Forest:**
    - **Concept:** It is easier to "isolate" an anomaly than a normal point.
    - **Process:** Randomly partitions data using trees. Anomalous points have shorter path lengths (they are isolated in fewer splits).

### Support Vector Machines (SVM)

- **One-Class SVM:** Learns a soft boundary that encompasses the bulk of the "normal" data. Anything outside this boundary is classified as an anomaly.

### Deep Learning Methods

- **Autoencoders:**
    - **Training:** The network is trained to compress and reconstruct normal data.
    - **Detection:** When an anomaly is passed through, the "Reconstruction Error" is high because the model has not learned to represent that specific pattern.

## D. Common Applications

- **Fraud Detection:** Identifying stolen credit cards or insurance fraud.
- **Intrusion Detection:** Monitoring network traffic for hacking or malware.
- **Industrial Monitoring:** Detecting equipment failure or sensor malfunctions.
- **Medical Diagnosis:** Identifying unusual patterns in ECG or MRI scans.

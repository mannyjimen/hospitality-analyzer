# Hospitality Analyzer

No one likes bad hospitality.  
What if there was a way to detect whether a business doesn't provide the best service before diving into the reviews?  
The goal of hospitality-analyzer is to gather reviews on thousands of businesses, and to figure out some signs that a business provides bad hospitality.  
For this specific project, we will be using Yelp's open dataset.

## Dataset Mapping

This is a self-made mapping of the relationships between all provided .json files from Yelp's dataset.  
![dataset-mapping](/assets/dataset_mapping.png)

## Data Pipeline Architecture

### Review Filter:
Given an arbitrary amount of reviews per city, we will be filtering out the reviews that signify unfair treatment, bad hospitality, or any other signs of negative behavior.
**Input**: many review.json files  
**Output**: review.json files that signify negative treatment.  
### Business Extractor
Given the remaining reviews, extract all businesses related to those reviews. Since some businesses might have more reviews than others, give more “bias” or “weight” to businesses based on the number of reviews in input.
**Input**: review.json files  
**Output**: `business_id`s and weights associated.   
### Similarity Finder
Find similarities between the given businesses. Give “bias” or “weight” to similarities that are more apparent.  
**Input**: various “business_id”s and weights associated.  
**Output**: similarities, and strength rating for each similarity.  

## Diagram
![data-pipeline-architecture](/assets/data_pipeline_architecture.png)
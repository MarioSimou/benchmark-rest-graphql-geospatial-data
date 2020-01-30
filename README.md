# GraphQL && Mongodb for Geospatial Data


## GraphqlQuery

```
query {
  population {
    type
    name
    features {
      ...featureFragment
    }
  }
}

fragment geometryFragment on Geometry {
  type
  coordinates
}

fragment propertiesFragment on Properties {
  gmlId
  localId
  namespace
  versionId
  notCountedProportion
  localisedCharacterString
  endPosition
}

fragment featureFragment on Feature {
 type
 properties {
   ...propertiesFragment
 }
 geometry {
   ...geometryFragment
 }
}

```
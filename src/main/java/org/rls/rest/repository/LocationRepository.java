package org.rls.rest.repository;

import org.rls.rest.model.Location;
import org.springframework.data.repository.PagingAndSortingRepository;
import org.springframework.data.rest.core.annotation.RepositoryRestResource;

@RepositoryRestResource(exported = false)
public interface LocationRepository extends PagingAndSortingRepository<Location, Long> {

    Location findByObjectId(int objectId);
}

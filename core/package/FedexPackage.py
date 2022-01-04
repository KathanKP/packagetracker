"""
This is a child class of BasePackage and contains logic about dealing with a Fedex package
"""
from .BasePackage import BasePackage

class FedexPackage(BasePackage):

    """
    Parameters:
        tracking_number: tracking number for this package (String)
        service: Service which is delivering this package eg. USPS, Fedex etc. (String)
        status: Status of the package (String)
    """
    def __init__(self, tracking_number, service, status):
        super().__init__(tracking_number, service, status)

    def track_package(self):
        return

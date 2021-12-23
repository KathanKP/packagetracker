"""
This is the class for Fedex API which inherits from the Base API class
"""

from .BaseAPI import BaseAPI


class FedexAPI(BaseAPI):

    def __init__(self, tracking_number):
        BaseAPI.__init__(self, tracking_number)

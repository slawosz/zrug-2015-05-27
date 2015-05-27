class WorkersController < ApplicationController
  def trigger
    SimpleWorker.perform_async('bob', 5)
    redirect_to products_url
  end
end

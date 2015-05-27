class SimpleWorker
  include Sidekiq::Worker

  def perform(*args)
    10.times do
      p "working with #{args}"
      sleep 1
    end
  end
end
